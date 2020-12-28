// Code generated by piper's step-generator. DO NOT EDIT.

package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/SAP/jenkins-library/pkg/config"
	"github.com/SAP/jenkins-library/pkg/log"
	"github.com/SAP/jenkins-library/pkg/telemetry"
	"github.com/spf13/cobra"
)

type executeNewmanOptions struct {
	Verbose              bool   `json:"verbose,omitempty"`
	NewmanCollection     string `json:"newmanCollection,omitempty"`
	NewmanRunCommand     string `json:"newmanRunCommand,omitempty"`
	NewmanInstallCommand string `json:"newmanInstallCommand,omitempty"`
	FailOnError          bool   `json:"failOnError,omitempty"`
	CfAppsWithSecrets    bool   `json:"cfAppsWithSecrets,omitempty"`
}

// ExecuteNewmanCommand This script executes [Postman](https://www.getpostman.com) tests from a collection via the [Newman](https://www.getpostman.com/docs/v6/postman/collection_runs/command_line_integration_with_newman) command line tool.
func ExecuteNewmanCommand() *cobra.Command {
	const STEP_NAME = "executeNewman"

	metadata := executeNewmanMetadata()
	var stepConfig executeNewmanOptions
	var startTime time.Time

	var createExecuteNewmanCmd = &cobra.Command{
		Use:   STEP_NAME,
		Short: "This script executes [Postman](https://www.getpostman.com) tests from a collection via the [Newman](https://www.getpostman.com/docs/v6/postman/collection_runs/command_line_integration_with_newman) command line tool.",
		Long:  ``,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			startTime = time.Now()
			log.SetStepName(STEP_NAME)
			log.SetVerbose(GeneralConfig.Verbose)

			path, _ := os.Getwd()
			fatalHook := &log.FatalHook{CorrelationID: GeneralConfig.CorrelationID, Path: path}
			log.RegisterHook(fatalHook)

			err := PrepareConfig(cmd, &metadata, STEP_NAME, &stepConfig, config.OpenPiperFile)
			if err != nil {
				log.SetErrorCategory(log.ErrorConfiguration)
				return err
			}

			if len(GeneralConfig.HookConfig.SentryConfig.Dsn) > 0 {
				sentryHook := log.NewSentryHook(GeneralConfig.HookConfig.SentryConfig.Dsn, GeneralConfig.CorrelationID)
				log.RegisterHook(&sentryHook)
			}

			return nil
		},
		Run: func(_ *cobra.Command, _ []string) {
			telemetryData := telemetry.CustomData{}
			telemetryData.ErrorCode = "1"
			handler := func() {
				config.RemoveVaultSecretFiles()
				telemetryData.Duration = fmt.Sprintf("%v", time.Since(startTime).Milliseconds())
				telemetryData.ErrorCategory = log.GetErrorCategory().String()
				telemetry.Send(&telemetryData)
			}
			log.DeferExitHandler(handler)
			defer handler()
			telemetry.Initialize(GeneralConfig.NoTelemetry, STEP_NAME)
			executeNewman(stepConfig, &telemetryData)
			telemetryData.ErrorCode = "0"
			log.Entry().Info("SUCCESS")
		},
	}

	addExecuteNewmanFlags(createExecuteNewmanCmd, &stepConfig)
	return createExecuteNewmanCmd
}

func addExecuteNewmanFlags(cmd *cobra.Command, stepConfig *executeNewmanOptions) {
	cmd.Flags().BoolVar(&stepConfig.Verbose, "verbose", false, "Print more detailed information into the log.")
	cmd.Flags().StringVar(&stepConfig.NewmanCollection, "newmanCollection", os.Getenv("PIPER_newmanCollection"), "The test collection that should be executed. This could also be a file pattern.")
	cmd.Flags().StringVar(&stepConfig.NewmanRunCommand, "newmanRunCommand", os.Getenv("PIPER_newmanRunCommand"), "The newman command that will be executed inside the docker container.")
	cmd.Flags().StringVar(&stepConfig.NewmanInstallCommand, "newmanInstallCommand", os.Getenv("PIPER_newmanInstallCommand"), "The shell command that will be executed inside the docker container to install Newman.")
	cmd.Flags().BoolVar(&stepConfig.FailOnError, "failOnError", false, "Defines the behavior, in case tests fail.")
	cmd.Flags().BoolVar(&stepConfig.CfAppsWithSecrets, "cfAppsWithSecrets", false, "Define name array of cloud foundry apps deployed for which secrets (clientid and clientsecret) will be appended")

	cmd.MarkFlagRequired("verbose")
	cmd.MarkFlagRequired("newmanCollection")
	cmd.MarkFlagRequired("newmanRunCommand")
	cmd.MarkFlagRequired("newmanInstallCommand")
	cmd.MarkFlagRequired("failOnError")
	cmd.MarkFlagRequired("cfAppsWithSecrets")
}

// retrieve step metadata
func executeNewmanMetadata() config.StepData {
	var theMetaData = config.StepData{
		Metadata: config.StepMetadata{
			Name:        "executeNewman",
			Aliases:     []config.Alias{},
			Description: "This script executes [Postman](https://www.getpostman.com) tests from a collection via the [Newman](https://www.getpostman.com/docs/v6/postman/collection_runs/command_line_integration_with_newman) command line tool.",
		},
		Spec: config.StepSpec{
			Inputs: config.StepInputs{
				Parameters: []config.StepParameters{
					{
						Name:        "verbose",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "bool",
						Mandatory:   true,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "newmanCollection",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   true,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "newmanRunCommand",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   true,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "newmanInstallCommand",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   true,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "failOnError",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "bool",
						Mandatory:   true,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "cfAppsWithSecrets",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "bool",
						Mandatory:   true,
						Aliases:     []config.Alias{},
					},
				},
			},
		},
	}
	return theMetaData
}
