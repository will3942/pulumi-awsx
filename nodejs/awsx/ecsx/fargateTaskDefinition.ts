import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import * as role from "../role";
import * as utils from "../utils";
import { Container } from "./container";

export interface FargateTaskDefinitionArgs {
    // Properties copied from ecs.TaskDefinitionArgs
    /**
     * The vpc that the service for this task will run in.  Does not normally need to be explicitly
     * provided as it will be inferred from the cluster the service is associated with.
     */
    vpcId?: pulumi.Input<string>;

    /**
     * A set of placement constraints rules that are taken into consideration during task placement.
     * Maximum number of `placement_constraints` is `10`.
     */
    placementConstraints?: aws.ecs.TaskDefinitionArgs["placementConstraints"];

    /**
     * The proxy configuration details for the App Mesh proxy.
     */
    proxyConfiguration?: aws.ecs.TaskDefinitionArgs["proxyConfiguration"];

    /**
     * A set of volume blocks that containers in your task may use.
     */
    volumes?: aws.ecs.TaskDefinitionArgs["volumes"];

    // Properties we've added/changed.
    /**
     * Log group for logging information related to the service.  If `undefined` a default instance
     * with a one-day retention policy will be created.  If `null` no log group will be created.
     */
    logGroup?: aws.cloudwatch.LogGroup | aws.cloudwatch.LogGroupArgs;

    /**
     * IAM role that allows your Amazon ECS container task to make calls to other AWS services. If
     * `undefined`, a default will be created for the task.  If `null` no role will be created.
     */
    taskRole?: aws.iam.Role | RoleArgs;

    /**
     * An optional family name for the Task Definition. If not specified, then a suitable default will be created.
     */
    family?: pulumi.Input<string>;

    /**
     * The execution role that the Amazon ECS container agent and the Docker daemon can assume.
     *
     *  If `undefined`, a default will be created for the task.  If `null` no role will be created.
     */
    executionRole?: aws.iam.Role | RoleArgs;

    /**
     * The number of cpu units used by the task.  If not provided, a default will be computed
     * based on the cumulative needs specified by [containerDefinitions]
     */
    cpu?: pulumi.Input<string>;

    /**
     * The amount (in MiB) of memory used by the task.  If not provided, a default will be computed
     * based on the cumulative needs specified by [containerDefinitions]
     */
    memory?: pulumi.Input<string>;

    /**
     * Single container to make a TaskDefinition from.  Useful for simple cases where there aren't
     * multiple containers, especially when creating a TaskDefinition to call [run] on.
     *
     * Either [container] or [containers] must be provided.
     */
    container?: Container;

    /**
     * All the containers to make a TaskDefinition from.  Useful when creating a Service that will
     * contain many containers within.
     *
     * Either [container] or [containers] must be provided.
     */
    containers?: Record<string, Container>;

    /**
     * Key-value mapping of resource tags
     */
    tags?: pulumi.Input<aws.Tags>;
}

export interface RoleArgs {
    name?: string;
    assumeRolePolicy?: string | aws.iam.PolicyDocument;
    policyArns?: string[];
}

export class FargateTaskDefinition extends pulumi.ComponentResource {
    public readonly taskDefinition: aws.ecs.TaskDefinition;
    public readonly logGroup?: aws.cloudwatch.LogGroup;
    public readonly containers: Record<string, Container>;
    public readonly taskRole?: aws.iam.Role;
    public readonly executionRole?: aws.iam.Role;

    constructor(
        name: string,
        args: FargateTaskDefinitionArgs,
        opts: pulumi.ComponentResourceOptions = {}
    ) {
        if (
            (!args.container && !args.containers) ||
            (args.container && args.containers)
        ) {
            throw new Error(
                "One of [container] or [containers] must be provided"
            );
        }

        super("awsx:x:ecs:FargateTaskDefinition", name, {}, opts);

        if (args.logGroup) {
            this.logGroup = aws.cloudwatch.LogGroup.isInstance(args.logGroup)
                ? args.logGroup
                : new aws.cloudwatch.LogGroup(name, args.logGroup, opts);
        }

        if (args.taskRole) {
            this.taskRole = aws.iam.Role.isInstance(args.taskRole)
                ? args.taskRole
                : role.createRoleAndPolicies(
                      args.taskRole.name ?? `${name}-task`,
                      args.taskRole.assumeRolePolicy ??
                          defaultRoleAssumeRolePolicy(),
                      args.taskRole.policyArns ?? defaultTaskRolePolicyARNs(),
                      { ...opts, parent: this }
                  ).role;
        }

        if (args.executionRole) {
            this.executionRole = aws.iam.Role.isInstance(args.executionRole)
                ? args.executionRole
                : role.createRoleAndPolicies(
                      args.executionRole.name ?? `${name}-execution`,
                      args.executionRole.assumeRolePolicy ??
                          defaultRoleAssumeRolePolicy(),
                      args.executionRole.policyArns ??
                          defaultExecutionRolePolicyARNs(),
                      { ...opts, parent: this }
                  ).role;
        }

        this.containers = args.containers ?? { container: args.container };

        const containerDefinitions = computeContainerDefinitions(
            this,
            this.containers,
            this.logGroup?.id
        );

        this.taskDefinition = new aws.ecs.TaskDefinition(
            name,
            buildTaskDefinitionArgs(
                name,
                args,
                containerDefinitions,
                this.taskRole,
                this.executionRole
            ),
            { parent: this }
        );
    }
}

function buildTaskDefinitionArgs(
    name: string,
    args: FargateTaskDefinitionArgs,
    containerDefinitions: pulumi.Output<aws.ecs.ContainerDefinition[]>,
    taskRole?: aws.iam.Role,
    executionRole?: aws.iam.Role
): aws.ecs.TaskDefinitionArgs {
    const containerString = containerDefinitions.apply((d) =>
        JSON.stringify(d)
    );
    const defaultFamily = containerString.apply(
        (s) => name + "-" + utils.sha1hash(pulumi.getStack() + s)
    );
    const family = utils.ifUndefined(args.family, defaultFamily);

    return {
        ...args,
        family,
        containerDefinitions: containerString,
    };
}

function computeContainerDefinitions(
    parent: pulumi.Resource,
    containers: Record<string, Container>,
    logGroupId: pulumi.Input<string | undefined>
): pulumi.Output<aws.ecs.ContainerDefinition[]> {
    const result: pulumi.Output<aws.ecs.ContainerDefinition>[] = [];

    for (const containerName of Object.keys(containers)) {
        const container = containers[containerName];

        result.push(
            computeContainerDefinition(
                parent,
                containerName,
                container,
                logGroupId
            )
        );
    }

    return pulumi.all(result);
}

function computeContainerDefinition(
    parent: pulumi.Resource,
    containerName: string,
    container: Container,
    logGroupId: pulumi.Input<string | undefined>
): pulumi.Output<aws.ecs.ContainerDefinition> {
    const resolvedMappings = container.portMappings.map((mappingInput) =>
        aws.lb.TargetGroup.isInstance(mappingInput)
            ? mappingInput.port.apply(
                  (port): aws.ecs.PortMapping => ({
                      containerPort: port,
                      hostPort: port,
                  })
              )
            : mappingInput
    );
    const region = utils.getRegion(parent);
    return pulumi
        .all([container, resolvedMappings, region, logGroupId])
        .apply(([container, portMappings, region, logGroupId]) => {
            const containerDefinition: aws.ecs.ContainerDefinition = {
                ...container,
                portMappings,
                name: containerName,
            };
            if (
                containerDefinition.logConfiguration === undefined &&
                logGroupId !== undefined
            ) {
                containerDefinition.logConfiguration = {
                    logDriver: "awslogs",
                    options: {
                        "awslogs-group": logGroupId,
                        "awslogs-region": region,
                        "awslogs-stream-prefix": containerName,
                    },
                };
            }
            return containerDefinition;
        });
}

function defaultRoleAssumeRolePolicy(): aws.iam.PolicyDocument {
    return {
        Version: "2012-10-17",
        Statement: [
            {
                Action: "sts:AssumeRole",
                Principal: {
                    Service: "ecs-tasks.amazonaws.com",
                },
                Effect: "Allow",
                Sid: "",
            },
        ],
    };
}

function defaultTaskRolePolicyARNs() {
    return [
        // Provides full access to Lambda
        aws.iam.ManagedPolicy.LambdaFullAccess,
        // Required for lambda compute to be able to run Tasks
        aws.iam.ManagedPolicy.AmazonECSFullAccess,
    ];
}

function defaultExecutionRolePolicyARNs() {
    return [
        "arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy",
        aws.iam.ManagedPolicies.AWSLambdaBasicExecutionRole,
    ];
}
