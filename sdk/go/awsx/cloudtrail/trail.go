// *** WARNING: this file was generated by pulumi-gen-awsx. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudtrail

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/cloudtrail"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Trail struct {
	pulumi.ResourceState

	// ARN of the trail.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Region in which the trail was created.
	HomeRegion pulumi.StringPtrOutput `pulumi:"homeRegion"`
	// Map of tags to assign to the trail. If configured with provider defaultTags present, tags with matching keys will overwrite those defined at the provider-level.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewTrail registers a new resource with the given unique name, arguments, and options.
func NewTrail(ctx *pulumi.Context,
	name string, args *TrailArgs, opts ...pulumi.ResourceOption) (*Trail, error) {
	if args == nil {
		args = &TrailArgs{}
	}

	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("aws:cloudtrail:x:Trail"),
		},
	})
	opts = append(opts, aliases)
	var resource Trail
	err := ctx.RegisterRemoteComponentResource("awsx:cloudtrail:Trail", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type trailArgs struct {
	// Specifies an advanced event selector for enabling data event logging.
	AdvancedEventSelectors []cloudtrail.TrailAdvancedEventSelector `pulumi:"advancedEventSelectors"`
	// If sendToCloudWatchLogs is enabled, provide the log group configuration.
	CloudWatchLogGroupArgs *LogGroup `pulumi:"cloudWatchLogGroupArgs"`
	// Log group name using an ARN that represents the log group to which CloudTrail logs will be delivered. Note that CloudTrail requires the Log Stream wildcard.
	CloudWatchLogsGroupArn *string `pulumi:"cloudWatchLogsGroupArn"`
	// Role for the CloudWatch Logs endpoint to assume to write to a user’s log group.
	CloudWatchLogsRoleArn *string `pulumi:"cloudWatchLogsRoleArn"`
	// Whether log file integrity validation is enabled. Defaults to `false`.
	EnableLogFileValidation *bool `pulumi:"enableLogFileValidation"`
	// Enables logging for the trail. Defaults to `true`. Setting this to `false` will pause logging.
	EnableLogging *bool `pulumi:"enableLogging"`
	// Specifies an event selector for enabling data event logging. Please note the CloudTrail limits when configuring these
	EventSelectors []cloudtrail.TrailEventSelector `pulumi:"eventSelectors"`
	// Whether the trail is publishing events from global services such as IAM to the log files. Defaults to `true`.
	IncludeGlobalServiceEvents *bool `pulumi:"includeGlobalServiceEvents"`
	// Configuration block for identifying unusual operational activity.
	InsightSelectors []cloudtrail.TrailInsightSelector `pulumi:"insightSelectors"`
	// Whether the trail is created in the current region or in all regions. Defaults to `false`.
	IsMultiRegionTrail *bool `pulumi:"isMultiRegionTrail"`
	// Whether the trail is an AWS Organizations trail. Organization trails log events for the master account and all member accounts. Can only be created in the organization master account. Defaults to `false`
	IsOrganizationTrail *bool `pulumi:"isOrganizationTrail"`
	// KMS key ARN to use to encrypt the logs delivered by CloudTrail.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// Name of the S3 bucket designated for publishing log files.
	S3BucketName *string `pulumi:"s3BucketName"`
	// S3 key prefix that follows the name of the bucket you have designated for log file delivery.
	S3KeyPrefix *string `pulumi:"s3KeyPrefix"`
	// If CloudTrail pushes logs to CloudWatch Logs in addition to S3. Disabled by default to reduce costs. Defaults to `false`
	SendToCloudWatchLogs *bool `pulumi:"sendToCloudWatchLogs"`
	// Name of the Amazon SNS topic defined for notification of log file delivery.
	SnsTopicName *string `pulumi:"snsTopicName"`
	// Map of tags to assign to the trail. If configured with provider defaultTags present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Trail resource.
type TrailArgs struct {
	// Specifies an advanced event selector for enabling data event logging.
	AdvancedEventSelectors cloudtrail.TrailAdvancedEventSelectorArrayInput
	// If sendToCloudWatchLogs is enabled, provide the log group configuration.
	CloudWatchLogGroupArgs LogGroupPtrInput
	// Log group name using an ARN that represents the log group to which CloudTrail logs will be delivered. Note that CloudTrail requires the Log Stream wildcard.
	CloudWatchLogsGroupArn pulumi.StringPtrInput
	// Role for the CloudWatch Logs endpoint to assume to write to a user’s log group.
	CloudWatchLogsRoleArn pulumi.StringPtrInput
	// Whether log file integrity validation is enabled. Defaults to `false`.
	EnableLogFileValidation pulumi.BoolPtrInput
	// Enables logging for the trail. Defaults to `true`. Setting this to `false` will pause logging.
	EnableLogging pulumi.BoolPtrInput
	// Specifies an event selector for enabling data event logging. Please note the CloudTrail limits when configuring these
	EventSelectors cloudtrail.TrailEventSelectorArrayInput
	// Whether the trail is publishing events from global services such as IAM to the log files. Defaults to `true`.
	IncludeGlobalServiceEvents pulumi.BoolPtrInput
	// Configuration block for identifying unusual operational activity.
	InsightSelectors cloudtrail.TrailInsightSelectorArrayInput
	// Whether the trail is created in the current region or in all regions. Defaults to `false`.
	IsMultiRegionTrail pulumi.BoolPtrInput
	// Whether the trail is an AWS Organizations trail. Organization trails log events for the master account and all member accounts. Can only be created in the organization master account. Defaults to `false`
	IsOrganizationTrail pulumi.BoolPtrInput
	// KMS key ARN to use to encrypt the logs delivered by CloudTrail.
	KmsKeyId pulumi.StringPtrInput
	// Name of the S3 bucket designated for publishing log files.
	S3BucketName pulumi.StringPtrInput
	// S3 key prefix that follows the name of the bucket you have designated for log file delivery.
	S3KeyPrefix pulumi.StringPtrInput
	// If CloudTrail pushes logs to CloudWatch Logs in addition to S3. Disabled by default to reduce costs. Defaults to `false`
	SendToCloudWatchLogs pulumi.BoolPtrInput
	// Name of the Amazon SNS topic defined for notification of log file delivery.
	SnsTopicName pulumi.StringPtrInput
	// Map of tags to assign to the trail. If configured with provider defaultTags present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (TrailArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*trailArgs)(nil)).Elem()
}

type TrailInput interface {
	pulumi.Input

	ToTrailOutput() TrailOutput
	ToTrailOutputWithContext(ctx context.Context) TrailOutput
}

func (*Trail) ElementType() reflect.Type {
	return reflect.TypeOf((**Trail)(nil)).Elem()
}

func (i *Trail) ToTrailOutput() TrailOutput {
	return i.ToTrailOutputWithContext(context.Background())
}

func (i *Trail) ToTrailOutputWithContext(ctx context.Context) TrailOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrailOutput)
}

// TrailArrayInput is an input type that accepts TrailArray and TrailArrayOutput values.
// You can construct a concrete instance of `TrailArrayInput` via:
//
//          TrailArray{ TrailArgs{...} }
type TrailArrayInput interface {
	pulumi.Input

	ToTrailArrayOutput() TrailArrayOutput
	ToTrailArrayOutputWithContext(context.Context) TrailArrayOutput
}

type TrailArray []TrailInput

func (TrailArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Trail)(nil)).Elem()
}

func (i TrailArray) ToTrailArrayOutput() TrailArrayOutput {
	return i.ToTrailArrayOutputWithContext(context.Background())
}

func (i TrailArray) ToTrailArrayOutputWithContext(ctx context.Context) TrailArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrailArrayOutput)
}

// TrailMapInput is an input type that accepts TrailMap and TrailMapOutput values.
// You can construct a concrete instance of `TrailMapInput` via:
//
//          TrailMap{ "key": TrailArgs{...} }
type TrailMapInput interface {
	pulumi.Input

	ToTrailMapOutput() TrailMapOutput
	ToTrailMapOutputWithContext(context.Context) TrailMapOutput
}

type TrailMap map[string]TrailInput

func (TrailMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Trail)(nil)).Elem()
}

func (i TrailMap) ToTrailMapOutput() TrailMapOutput {
	return i.ToTrailMapOutputWithContext(context.Background())
}

func (i TrailMap) ToTrailMapOutputWithContext(ctx context.Context) TrailMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrailMapOutput)
}

type TrailOutput struct{ *pulumi.OutputState }

func (TrailOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Trail)(nil)).Elem()
}

func (o TrailOutput) ToTrailOutput() TrailOutput {
	return o
}

func (o TrailOutput) ToTrailOutputWithContext(ctx context.Context) TrailOutput {
	return o
}

type TrailArrayOutput struct{ *pulumi.OutputState }

func (TrailArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Trail)(nil)).Elem()
}

func (o TrailArrayOutput) ToTrailArrayOutput() TrailArrayOutput {
	return o
}

func (o TrailArrayOutput) ToTrailArrayOutputWithContext(ctx context.Context) TrailArrayOutput {
	return o
}

func (o TrailArrayOutput) Index(i pulumi.IntInput) TrailOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Trail {
		return vs[0].([]*Trail)[vs[1].(int)]
	}).(TrailOutput)
}

type TrailMapOutput struct{ *pulumi.OutputState }

func (TrailMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Trail)(nil)).Elem()
}

func (o TrailMapOutput) ToTrailMapOutput() TrailMapOutput {
	return o
}

func (o TrailMapOutput) ToTrailMapOutputWithContext(ctx context.Context) TrailMapOutput {
	return o
}

func (o TrailMapOutput) MapIndex(k pulumi.StringInput) TrailOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Trail {
		return vs[0].(map[string]*Trail)[vs[1].(string)]
	}).(TrailOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TrailInput)(nil)).Elem(), &Trail{})
	pulumi.RegisterInputType(reflect.TypeOf((*TrailArrayInput)(nil)).Elem(), TrailArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TrailMapInput)(nil)).Elem(), TrailMap{})
	pulumi.RegisterOutputType(TrailOutput{})
	pulumi.RegisterOutputType(TrailArrayOutput{})
	pulumi.RegisterOutputType(TrailMapOutput{})
}
