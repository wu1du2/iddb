/*
   Copyright 2014 CoreOS, Inc.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package flags

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/coreos/etcd/pkg/transport"
)

// DeprecatedFlag encapsulates a flag that may have been previously valid but
// is now deprecated. If a DeprecatedFlag is set, an error occurs.
type DeprecatedFlag struct {
	Name string
}

func (f *DeprecatedFlag) Set(_ string) error {
	return fmt.Errorf(`flag "-%s" is no longer supported.`, f.Name)
}

func (f *DeprecatedFlag) String() string {
	return ""
}

// IgnoredFlag encapsulates a flag that may have been previously valid but is
// now ignored. If an IgnoredFlag is set, a warning is printed and
// operation continues.
type IgnoredFlag struct {
	Name string
}

// IsBoolFlag is defined to allow the flag to be defined without an argument
func (f *IgnoredFlag) IsBoolFlag() bool {
	return true
}

func (f *IgnoredFlag) Set(s string) error {
	log.Printf(`WARNING: flag "-%s" is no longer supported - ignoring.`, f.Name)
	return nil
}

func (f *IgnoredFlag) String() string {
	return ""
}

// SetFlagsFromEnv parses all registered flags in the given flagset,
// and if they are not already set it attempts to set their values from
// environment variables. Environment variables take the name of the flag but
// are UPPERCASE, have the prefix "ETCD_", and any dashes are replaced by
// underscores - for example: some-flag => ETCD_SOME_FLAG
func SetFlagsFromEnv(fs *flag.FlagSet) error {
	var err error
	alreadySet := make(map[string]bool)
	fs.Visit(func(f *flag.Flag) {
		alreadySet[f.Name] = true
	})
	fs.VisitAll(func(f *flag.Flag) {
		if !alreadySet[f.Name] {
			key := "ETCD_" + strings.ToUpper(strings.Replace(f.Name, "-", "_", -1))
			val := os.Getenv(key)
			if val != "" {
				if serr := fs.Set(f.Name, val); serr != nil {
					err = fmt.Errorf("invalid value %q for %s: %v", val, key, serr)
				}
			}
		}
	})
	return err
}

// URLsFromFlags decides what URLs should be using two different flags
// as datasources. The first flag's Value must be of type URLs, while
// the second must be of type IPAddressPort. If both of these flags
// are set, an error will be returned. If only the first flag is set,
// the underlying url.URL objects will be returned unmodified. If the
// second flag happens to be set, the underlying IPAddressPort will be
// converted to a url.URL and returned. The Scheme of the returned
// url.URL will be http unless the provided TLSInfo object is non-empty.
// If neither of the flags have been explicitly set, the default value
// of the first flag will be returned unmodified.
func URLsFromFlags(fs *flag.FlagSet, urlsFlagName string, addrFlagName string, tlsInfo transport.TLSInfo) ([]url.URL, error) {
	visited := make(map[string]struct{})
	fs.Visit(func(f *flag.Flag) {
		visited[f.Name] = struct{}{}
	})

	_, urlsFlagIsSet := visited[urlsFlagName]
	_, addrFlagIsSet := visited[addrFlagName]

	if addrFlagIsSet {
		if urlsFlagIsSet {
			return nil, fmt.Errorf("Set only one of flags -%s and -%s", urlsFlagName, addrFlagName)
		}

		addr := *fs.Lookup(addrFlagName).Value.(*IPAddressPort)
		addrURL := url.URL{Scheme: "http", Host: addr.String()}
		if !tlsInfo.Empty() {
			addrURL.Scheme = "https"
		}
		return []url.URL{addrURL}, nil
	}

	return []url.URL(*fs.Lookup(urlsFlagName).Value.(*URLsValue)), nil
}
