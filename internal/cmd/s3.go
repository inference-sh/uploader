package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/inference-sh/uploader/internal/upload"
)

func newS3Cmd(cfg *upload.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "s3",
		Short: "Upload a file to S3-compatible storage",
		PreRun: func(cmd *cobra.Command, args []string) {
			// Env var fallbacks — flags take precedence
			envDefault(cmd, "endpoint", "S3_ENDPOINT")
			envDefault(cmd, "region", "S3_REGION")
			envDefault(cmd, "access-key-id", "S3_ACCESS_KEY_ID")
			envDefault(cmd, "access-key-secret", "S3_SECRET_ACCESS_KEY")
			envDefault(cmd, "bucket", "S3_BUCKET")
			envDefault(cmd, "url", "S3_URL")
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return upload.Upload(cfg)
		},
	}

	cmd.Flags().StringVar(&cfg.SourcePath, "source", "", "Source file path to upload")
	cmd.Flags().StringVar(&cfg.DestPath, "dest", "", "Destination path in S3")
	cmd.Flags().StringVar(&cfg.Endpoint, "endpoint", "", "S3 endpoint (env: S3_ENDPOINT)")
	cmd.Flags().StringVar(&cfg.Region, "region", "", "S3 region (env: S3_REGION)")
	cmd.Flags().StringVar(&cfg.AccessKeyID, "access-key-id", "", "S3 access key ID (env: S3_ACCESS_KEY_ID)")
	cmd.Flags().StringVar(&cfg.AccessKeySecret, "access-key-secret", "", "S3 access key secret (env: S3_SECRET_ACCESS_KEY)")
	cmd.Flags().StringVar(&cfg.Bucket, "bucket", "", "S3 bucket name (env: S3_BUCKET)")
	cmd.Flags().StringVar(&cfg.URL, "url", "", "Public base URL for the bucket (env: S3_URL)")

	cmd.MarkFlagRequired("source")
	cmd.MarkFlagRequired("dest")

	return cmd
}

// envDefault sets a flag's value from an env var if the flag wasn't explicitly set.
func envDefault(cmd *cobra.Command, flagName, envVar string) {
	f := cmd.Flags().Lookup(flagName)
	if f == nil || f.Changed {
		return
	}
	if v := os.Getenv(envVar); v != "" {
		f.Value.Set(v)
	}
}
