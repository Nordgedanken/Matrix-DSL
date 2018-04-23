// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"errors"
	"fmt"
	"github.com/Nordgedanken/matrix_dsl/cmd/generator/js"
	"github.com/Nordgedanken/matrix_dsl/cmd/lexer"
	"github.com/alecthomas/participle"
	"github.com/spf13/cobra"
	"os"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires at least one arg")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		parser, err := participle.Build(&lexer.Matrix{}, nil)
		if err != nil {
			return err
		}

		mx := &lexer.Matrix{}
		file, err := os.Open(args[0])
		if err != nil {
			return err
		}

		err = parser.Parse(file, mx)
		if err != nil {
			return err
		}

		printMX(mx)

		for _, v := range mx.Sections {
			if v.Identifier == "BOT" {
				js.GenerateBot(mx.Sections[0])
			} else {
				return errors.New("unknown Section type")
			}
		}
		return nil
	},
}

//printMX is a temporary pretty print of the generated AST
func printMX(mx *lexer.Matrix) error {
	if mx.Properties != nil {
	} else {
		fmt.Println("No Properties found")
	}
	if mx.Sections != nil {
		for _, s := range mx.Sections {
			fmt.Println("Section Name: ", s.Identifier)
			if s.Properties != nil {
				for _, p := range s.Properties {
					fmt.Printf("[%s] Prop Key: %s\n", s.Identifier, p.Key)
					if p.Event != nil {
						fmt.Printf("[%s] Event: %s\n", s.Identifier, *p.Event)
					}
					if p.Value != nil && p.Value.String != nil {
						fmt.Printf("[%s] Value: %s\n", s.Identifier, *p.Value.String)
					}
					if p.Arrays != nil {
						for i, a := range p.Arrays {
							fmt.Printf("[%s][%s][%d] Array Index: %d\n", s.Identifier, p.Key, i, i)
							if a.Key != "" {
								return errors.New("array key should always be empty")
							}
							if a.Properties != nil {
								for _, ap := range a.Properties {
									fmt.Printf("[%s][%s][%d] Prop Key: %s\n", s.Identifier, p.Key, i, ap.Key)
									if ap.Event != nil {
										fmt.Printf("[%s][%s][%d] Event: %s\n", s.Identifier, p.Key, i, *ap.Event)
									}
									if ap.Value != nil && ap.Value.String != nil {
										fmt.Printf("[%s][%s][%d] Value: %s\n", s.Identifier, p.Key, i, *ap.Value.String)
									}
								}
							} else {
								fmt.Printf("[%s][%s] No Arrays found\n", s.Identifier, p.Key)
							}
						}
					} else {
						fmt.Printf("[%s][%s] No Arrays found\n", s.Identifier, p.Key)
					}
				}
			} else {
				fmt.Printf("[%s] No Properties found\n", s.Identifier)
			}
		}
	} else {
		fmt.Println("No Sections found")
	}
	return nil
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
