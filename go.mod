module github.com/antonybholmes/go-mailserver

go 1.27

replace github.com/antonybholmes/go-sys => ../go-sys

require (
	github.com/antonybholmes/go-sys v0.0.0-20261110151334-69bf4fef33df
	github.com/aws/aws-sdk-go-v2 v1.48.0
	github.com/aws/aws-sdk-go-v2/config v1.33.18
	github.com/aws/aws-sdk-go-v2/service/sesv2 v1.79.9
	github.com/aws/aws-sdk-go-v2/service/sqs v1.53.0
	github.com/redis/go-redis/v9 v9.23.0
	github.com/segmentio/kafka-go v0.4.51
)

require (
	github.com/aws/aws-sdk-go-v2/credentials v1.21.0 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.20.9 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.5.12 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.8.12 // indirect
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.5.12 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.13.25 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.14.12 // indirect
	github.com/aws/aws-sdk-go-v2/service/signin v1.13.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.39.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.44.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/sts v1.52.0 // indirect
	github.com/aws/smithy-go v1.28.9 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/klauspost/compress v1.20.4 // indirect
	github.com/mattn/go-colorable v0.1.16 // indirect
	github.com/mattn/go-isatty v0.0.24 // indirect
	github.com/pierrec/lz4/v4 v4.1.33 // indirect
	github.com/rs/zerolog v1.35.1 // indirect
	go.uber.org/atomic v1.12.0 // indirect
	golang.org/x/sys v0.50.0 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.2.1 // indirect
)
