# Electric Gopher

An experimental Go library for interacting with Tesla's undocumented vehicle owner API.

# Notes

* This is experimental.
* The APIs are not supported and not documented, so this is mostly the result of various web resources and reverse engineering. If you are trying to use this library and run across any issues, differences or have any interesting notes, please log an issue or submit a PR.
* I'm relatively new to the Go ecosystem, so if you find the source to be awkward or non-idiomatic, please log an issue with any suggestions and reference links.
* No tests ... yet.
* Thanks for checking the project out!

# Usage

    // construct a client
    client = api.NewClient(
      "CLIENT ID",
      "CLIENT SECRET",
      "OWNER EMAIL",
      "OWNER PASSWORD",
      "https://owner-api.teslamotors.com",
      nil,
    )

    // authenticate
    err := client.Authenticate()
    if err != nil {
      log.Fatalf("Error ... %v", err)
    }

    // get some vehicles
    gvo, err := client.GetVehicles()
    if err != nil {
      log.Fatalf("Error ... %v", err)
    }
    log.Printf("Vehicles ... %v", gvo)

# Development

This project is currently developed against Go v1.10 and uses `dep` for dependency management and `make` for build functions.

# License

Copyright 2018 Nathan Beyer

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License. You may obtain a copy of the License at

&nbsp;&nbsp;&nbsp;&nbsp;http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.
