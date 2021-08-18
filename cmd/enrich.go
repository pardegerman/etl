/*
Copyright © 2021 Pär Degerman <par@degerman.org>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"os"
	"sort"

	"github.com/pardegerman/etl/singer"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func contains(s []string, searchterm string) bool {
	i := sort.SearchStrings(s, searchterm)
	return i < len(s) && s[i] == searchterm
}

// enrichCmd represents the catalog command
var enrichCmd = &cobra.Command{
	Use:   "enrich <DATABASE>",
	Short: "Enrich the catalog provided by a tap ran in discovery mode",
	Long: `When a tap is ran in discovery mode it provides a list of
streams, a catalog. Using this action, the information
in the configuration file is used to enrich the catalog
by adding non discoverable items. Feed the original
catalog on STDIN and the enriched catalog will be provided
on STDOUT.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cfg := viper.Sub("enrich")
		replicationMethod := cfg.GetString("replication-method")
		replicationKey := cfg.GetString("replication-key")
		excludeDatabases := cfg.GetStringSlice("exclude-databases")
		excludeTables := cfg.GetStringSlice("exclude-tables")

		// Prepare to use the contains for slices by sorting the exclude-* slices
		sort.Strings(excludeDatabases)
		sort.Strings(excludeTables)
		selected := true

		dbName := args[0]
		if contains(excludeDatabases, dbName) {
			err := fmt.Errorf("Database %s is marked for exclusion, nothing will be selected", dbName)
			cobra.CheckErr(err)
		}

		// Read the singer catalog from STDIN
		catalog, err := singer.ReadCatalog(os.Stdin)
		cobra.CheckErr(err)

		for _, stream := range catalog.Streams {
			if !contains(excludeTables, stream.TableName) {
				for _, metaData := range stream.Metadata {
					if len(metaData.Breadcrumb) == 0 && metaData.MetadataProps.DatabaseName == dbName {
						metaData.MetadataProps.Selected = &selected
						metaData.MetadataProps.ReplicationMethod = replicationMethod
						metaData.MetadataProps.ReplicationKey = replicationKey
						break
					}
				}
			}
		}

		out, err := catalog.Dump()
		cobra.CheckErr(err)

		fmt.Fprint(os.Stdout, out)
	},
}

func init() {
	rootCmd.AddCommand(enrichCmd)

	enrichCmd.PersistentFlags().String("replication-method", "LOG_BASED", "The default replication method to use")
	viper.BindPFlag("enrich.replication-method", enrichCmd.Flags().Lookup("replication-method"))
	viper.SetDefault("enrich.replication-method", "LOG_BASED")

	enrichCmd.PersistentFlags().String("replication-key", "", "The default replication key to use")
	viper.BindPFlag("enrich.replication-key", enrichCmd.Flags().Lookup("replication-key"))

	enrichCmd.PersistentFlags().StringSlice("exclude-databases", nil, "List of databases to exclude from replication")
	viper.BindPFlag("enrich.exclude-databases", enrichCmd.Flags().Lookup("exclude-databases"))

	enrichCmd.PersistentFlags().StringSlice("exclude-tables", nil, "List of tables to exclude from replication")
	viper.BindPFlag("enrich.exclude-tables", enrichCmd.Flags().Lookup("exclude-tables"))

}
