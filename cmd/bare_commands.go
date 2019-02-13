package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/mattn/go-colorable"
	gitcmd "github.com/tsuyoshiwada/go-gitcmd"
)

// ShimConf : Conf loader that stores the variables needed to run the library without URFave CLI
type ShimConf struct {
	Init       bool
	ConfigPath string
	OutputPath string
	Query      string
	NextTag    string
	Silent     bool
	NoColor    bool
	NoEmoji    bool
	SemVerOnly bool
}

// Do : Start up the git-chnglog commands without using the Urfave CLI
func (c *ShimConf) Do() (err error) {
	// Check for what we need
	err = c.Check()
	if err != nil {
		return err
	}

	// Set Generator if SemVer is requested via bool
	g := NewGenerator()
	if c.SemVerOnly {
		g = NewGeneratorSemVer()
	}

	var ex int
	wd, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to get working directory", err)
		return err
	}

	// initializer
	if c.Init {
		initializer := NewInitializer(
			&InitContext{
				WorkingDir: wd,
				Stdout:     colorable.NewColorableStdout(),
				Stderr:     colorable.NewColorableStderr(),
			},
			fs,
			NewQuestioner(
				gitcmd.New(&gitcmd.Config{
					Bin: "git",
				}),
				fs,
			),
			NewConfigBuilder(),
			templateBuilderFactory,
		)

		ex = initializer.Run()
		if ex != 0 {
			err = errors.New("I tried to run you through the init options, but it didn't work at all")
			return err
		}
	}

	// chglog initialize the CLI
	chglogCLI := NewCLI(
		&CLIContext{
			WorkingDir: wd,
			Stdout:     colorable.NewColorableStdout(),
			Stderr:     colorable.NewColorableStderr(),
			ConfigPath: c.ConfigPath,
			OutputPath: c.OutputPath,
			Silent:     c.Silent,
			NoColor:    c.NoColor,
			NoEmoji:    c.NoEmoji,
			Query:      c.Query,
			NextTag:    c.NextTag,
		},
		fs,
		NewConfigLoader(),
		g,
	)

	ex = chglogCLI.Run()
	if ex != 0 {
		err = errors.New("I tried to run what you asked, but it didn't work at all")
		return err
	}

	return nil
}

// Check : Make sure all the needed string configs are set
func (c *ShimConf) Check() (err error) {
	s := []string{
		"ConfigPath",
		"Query",
		"NextTag",
	}

	z := []string{
		c.ConfigPath,
		c.Query,
		c.NextTag,
	}

	for i, x := range z {
		if x == "" {
			t := "Missing required config in Config struct: " + s[i]
			err = errors.New(t)
			return err
		}
	}

	return nil
}
