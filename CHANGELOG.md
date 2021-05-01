# Changelog

## [v1.0.6](https://github.com/infobloxopen/terraform-provider-infoblox/tree/v1.0.6) (2021-04-23)

[Full Changelog](https://github.com/infobloxopen/terraform-provider-infoblox/compare/v1.1.1...v1.0.6)

## [v1.1.1](https://github.com/infobloxopen/terraform-provider-infoblox/tree/v1.1.1) (2021-04-23)

[Full Changelog](https://github.com/infobloxopen/terraform-provider-infoblox/compare/v1.1.0...v1.1.1)

**Implemented enhancements:**

- Resources are not truly destroyed upon destroy [\#57](https://github.com/infobloxopen/terraform-provider-infoblox/issues/57)
- Can't dynamically request a network [\#27](https://github.com/infobloxopen/terraform-provider-infoblox/issues/27)

**Closed issues:**

- Feature request: ability to change extended attributes [\#77](https://github.com/infobloxopen/terraform-provider-infoblox/issues/77)
- Commit changes? [\#69](https://github.com/infobloxopen/terraform-provider-infoblox/issues/69)
- DHCP Host records won't store in state [\#68](https://github.com/infobloxopen/terraform-provider-infoblox/issues/68)
- Terraform Provider Development Program - Second Review [\#66](https://github.com/infobloxopen/terraform-provider-infoblox/issues/66)
- Passing credentials without using environment variables [\#65](https://github.com/infobloxopen/terraform-provider-infoblox/issues/65)
- inconsistent vendoring [\#59](https://github.com/infobloxopen/terraform-provider-infoblox/issues/59)
- make build fails [\#34](https://github.com/infobloxopen/terraform-provider-infoblox/issues/34)
- Error accessing infoblox thru https proxy [\#28](https://github.com/infobloxopen/terraform-provider-infoblox/issues/28)

**Merged pull requests:**

- Remove old release workflow [\#111](https://github.com/infobloxopen/terraform-provider-infoblox/pull/111) ([cgroschupp](https://github.com/cgroschupp))
- Display travis build status of master and develop branch [\#109](https://github.com/infobloxopen/terraform-provider-infoblox/pull/109) ([somashekhar](https://github.com/somashekhar))
- Add provider to terraform registry [\#107](https://github.com/infobloxopen/terraform-provider-infoblox/pull/107) ([cgroschupp](https://github.com/cgroschupp))
- Revert "doc change" [\#91](https://github.com/infobloxopen/terraform-provider-infoblox/pull/91) ([somashekhar](https://github.com/somashekhar))
- doc change [\#90](https://github.com/infobloxopen/terraform-provider-infoblox/pull/90) ([somashekhar](https://github.com/somashekhar))
- doc change [\#85](https://github.com/infobloxopen/terraform-provider-infoblox/pull/85) ([somashekhar](https://github.com/somashekhar))
- Example NIOS and AWS/VMWare/Azure tf files to support Next Available Network and DNS/DHCP Records support for AWS. [\#84](https://github.com/infobloxopen/terraform-provider-infoblox/pull/84) ([somashekhar](https://github.com/somashekhar))
- Check that parent exists before network allocation. [\#83](https://github.com/infobloxopen/terraform-provider-infoblox/pull/83) ([AliaksandrDziarkach](https://github.com/AliaksandrDziarkach))
- Migration from terraform providers to our own repository [\#78](https://github.com/infobloxopen/terraform-provider-infoblox/pull/78) ([AvRajath](https://github.com/AvRajath))

## [v1.1.0](https://github.com/infobloxopen/terraform-provider-infoblox/tree/v1.1.0) (2020-08-11)

[Full Changelog](https://github.com/infobloxopen/terraform-provider-infoblox/compare/v1.0.5...v1.1.0)

**Closed issues:**

- Terraform provider documentation is not accurate on terraform.io website [\#76](https://github.com/infobloxopen/terraform-provider-infoblox/issues/76)
- change the package in main .go [\#75](https://github.com/infobloxopen/terraform-provider-infoblox/issues/75)

## [v1.0.5](https://github.com/infobloxopen/terraform-provider-infoblox/tree/v1.0.5) (2020-05-15)

[Full Changelog](https://github.com/infobloxopen/terraform-provider-infoblox/compare/v1.0.4...v1.0.5)

## [v1.0.4](https://github.com/infobloxopen/terraform-provider-infoblox/tree/v1.0.4) (2020-05-15)

[Full Changelog](https://github.com/infobloxopen/terraform-provider-infoblox/compare/v1.0.3...v1.0.4)

## [v1.0.3](https://github.com/infobloxopen/terraform-provider-infoblox/tree/v1.0.3) (2020-05-15)

[Full Changelog](https://github.com/infobloxopen/terraform-provider-infoblox/compare/v1.0.2...v1.0.3)

**Closed issues:**

- Are you able to make releases through CI? [\#33](https://github.com/infobloxopen/terraform-provider-infoblox/issues/33)

## [v1.0.2](https://github.com/infobloxopen/terraform-provider-infoblox/tree/v1.0.2) (2020-05-14)

[Full Changelog](https://github.com/infobloxopen/terraform-provider-infoblox/compare/v1.0.1...v1.0.2)

**Closed issues:**

- a\_record generated with duplicate domain name [\#71](https://github.com/infobloxopen/terraform-provider-infoblox/issues/71)
- Error creating A Record from network block [\#62](https://github.com/infobloxopen/terraform-provider-infoblox/issues/62)
- Too many errors [\#61](https://github.com/infobloxopen/terraform-provider-infoblox/issues/61)
- Resource: infoblox\_cname\_record by default appending zone information to alias record [\#60](https://github.com/infobloxopen/terraform-provider-infoblox/issues/60)
- Improved documentation [\#55](https://github.com/infobloxopen/terraform-provider-infoblox/issues/55)
- terraform doesn't change a HostRecord name when the vm\_name is changed in the block, only the EA vm\_name is changed, not the actual name [\#54](https://github.com/infobloxopen/terraform-provider-infoblox/issues/54)
- Add HostRecord function [\#53](https://github.com/infobloxopen/terraform-provider-infoblox/issues/53)
- terraform doesn't like it when a managed IP gets removed from infoblox by hand [\#52](https://github.com/infobloxopen/terraform-provider-infoblox/issues/52)
- infoblox\_ip\_association causing crash [\#49](https://github.com/infobloxopen/terraform-provider-infoblox/issues/49)
- infoblox\_ip\_allocation fails - doesn't know what to do with vm\_name field [\#48](https://github.com/infobloxopen/terraform-provider-infoblox/issues/48)
- Terraform Provider Development Program - Review [\#45](https://github.com/infobloxopen/terraform-provider-infoblox/issues/45)
- "stock" go make build fails [\#43](https://github.com/infobloxopen/terraform-provider-infoblox/issues/43)
- API error [\#32](https://github.com/infobloxopen/terraform-provider-infoblox/issues/32)
- Official terraform provider status [\#30](https://github.com/infobloxopen/terraform-provider-infoblox/issues/30)
- Does this provider work with Terraform 0.12.1 [\#29](https://github.com/infobloxopen/terraform-provider-infoblox/issues/29)
- Tenant ID Parameter? [\#21](https://github.com/infobloxopen/terraform-provider-infoblox/issues/21)
- Error using your Terraform provider to access Infoblox api [\#20](https://github.com/infobloxopen/terraform-provider-infoblox/issues/20)

**Merged pull requests:**

- New datasource infoblox network [\#74](https://github.com/infobloxopen/terraform-provider-infoblox/pull/74) ([pearcec](https://github.com/pearcec))
- Delete Network View fix issue \#57 [\#73](https://github.com/infobloxopen/terraform-provider-infoblox/pull/73) ([pearcec](https://github.com/pearcec))
- Leverage Github Actions for Building, Tagging, and Releasing the provider [\#70](https://github.com/infobloxopen/terraform-provider-infoblox/pull/70) ([NickLarsenNZ](https://github.com/NickLarsenNZ))
- Modify CNAME resource create function to remove zone [\#64](https://github.com/infobloxopen/terraform-provider-infoblox/pull/64) ([elangoganesan](https://github.com/elangoganesan))
- Updated Readme for use case without CNA license [\#63](https://github.com/infobloxopen/terraform-provider-infoblox/pull/63) ([AvRajath](https://github.com/AvRajath))
- Terraform review program [\#47](https://github.com/infobloxopen/terraform-provider-infoblox/pull/47) ([saiprasannasastry](https://github.com/saiprasannasastry))
- merge conflicts [\#46](https://github.com/infobloxopen/terraform-provider-infoblox/pull/46) ([saiprasannasastry](https://github.com/saiprasannasastry))
- Moving to go mod and bumping terraform version [\#42](https://github.com/infobloxopen/terraform-provider-infoblox/pull/42) ([saiprasannasastry](https://github.com/saiprasannasastry))

## [v1.0.1](https://github.com/infobloxopen/terraform-provider-infoblox/tree/v1.0.1) (2019-09-05)

[Full Changelog](https://github.com/infobloxopen/terraform-provider-infoblox/compare/v1.0.0...v1.0.1)

**Merged pull requests:**

- Fix build [\#39](https://github.com/infobloxopen/terraform-provider-infoblox/pull/39) ([mikecook](https://github.com/mikecook))
- Dsbrng25b remove license check [\#38](https://github.com/infobloxopen/terraform-provider-infoblox/pull/38) ([saiprasannasastry](https://github.com/saiprasannasastry))
- Addressing some internal QA bugs [\#37](https://github.com/infobloxopen/terraform-provider-infoblox/pull/37) ([saiprasannasastry](https://github.com/saiprasannasastry))
- Remove license check entirely [\#36](https://github.com/infobloxopen/terraform-provider-infoblox/pull/36) ([dvob](https://github.com/dvob))

## [v1.0.0](https://github.com/infobloxopen/terraform-provider-infoblox/tree/v1.0.0) (2019-08-09)

[Full Changelog](https://github.com/infobloxopen/terraform-provider-infoblox/compare/ec354d7410947945fe01f876269ded480f18029d...v1.0.0)

**Implemented enhancements:**

- Support for CNAME or A Records [\#22](https://github.com/infobloxopen/terraform-provider-infoblox/issues/22)

**Closed issues:**

- Error on plan [\#26](https://github.com/infobloxopen/terraform-provider-infoblox/issues/26)

**Merged pull requests:**

- CHANGELOG.md updated to v1.0.0 [\#35](https://github.com/infobloxopen/terraform-provider-infoblox/pull/35) ([jkraj](https://github.com/jkraj))
- Minor readability changes to README.md [\#31](https://github.com/infobloxopen/terraform-provider-infoblox/pull/31) ([scottsuch](https://github.com/scottsuch))
- This commit contains entire DNS changes [\#25](https://github.com/infobloxopen/terraform-provider-infoblox/pull/25) ([saiprasannasastry](https://github.com/saiprasannasastry))
- This commit contains patch to rebase issue [\#19](https://github.com/infobloxopen/terraform-provider-infoblox/pull/19) ([saiprasannasastry](https://github.com/saiprasannasastry))
- Update README.md [\#18](https://github.com/infobloxopen/terraform-provider-infoblox/pull/18) ([saiprasannasastry](https://github.com/saiprasannasastry))
- Example [\#17](https://github.com/infobloxopen/terraform-provider-infoblox/pull/17) ([saiprasannasastry](https://github.com/saiprasannasastry))
- Update README.md [\#16](https://github.com/infobloxopen/terraform-provider-infoblox/pull/16) ([saiprasannasastry](https://github.com/saiprasannasastry))
- Added block of code to reserve n number of IP's [\#14](https://github.com/infobloxopen/terraform-provider-infoblox/pull/14) ([saiprasannasastry](https://github.com/saiprasannasastry))
- Added an example tf for Vmware Vshpere [\#12](https://github.com/infobloxopen/terraform-provider-infoblox/pull/12) ([saiprasannasastry](https://github.com/saiprasannasastry))
- UT's for Infoblox-Provider [\#11](https://github.com/infobloxopen/terraform-provider-infoblox/pull/11) ([saiprasannasastry](https://github.com/saiprasannasastry))
- This commit contains a new resource IPAssociation [\#10](https://github.com/infobloxopen/terraform-provider-infoblox/pull/10) ([saiprasannasastry](https://github.com/saiprasannasastry))
- Validation of Cloud License [\#8](https://github.com/infobloxopen/terraform-provider-infoblox/pull/8) ([saiprasannasastry](https://github.com/saiprasannasastry))
- Validation of gateway ip existence [\#5](https://github.com/infobloxopen/terraform-provider-infoblox/pull/5) ([saiprasannasastry](https://github.com/saiprasannasastry))
- Modified makefile to build and gofmtcheck [\#3](https://github.com/infobloxopen/terraform-provider-infoblox/pull/3) ([jkraj](https://github.com/jkraj))
- Support for network and ip allocation [\#1](https://github.com/infobloxopen/terraform-provider-infoblox/pull/1) ([saiprasannasastry](https://github.com/saiprasannasastry))



\* *This Changelog was automatically generated       by [github_changelog_generator]      (https://github.com/github-changelog-generator/github-changelog-generator)*
