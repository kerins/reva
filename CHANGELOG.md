Changelog for reva 1.2.0 (2020-08-25)
=======================================

The following sections list the changes in reva 1.2.0 relevant to
reva users. The changes are ordered by importance.

Summary
-------

 * Fix #1099: Do not restore recycle entry on purge
 * Fix #1091: Allow listing the trashbin
 * Fix #1103: Restore and delete trash items via ocs
 * Fix #1090: Ensure ignoring public stray shares
 * Fix #1064: Ensure ignoring stray shares
 * Fix #1082: Minor fixes in reva cmd, gateway uploads and smtpclient
 * Fix #1115: Owncloud driver - propagate mtime on RemoveGrant
 * Fix #1111: Handle redirection prefixes when extracting destination from URL
 * Enh #1101: Add UID and GID in ldap auth driver
 * Enh #1077: Add calens check to verify changelog entries in CI
 * Enh #1072: Refactor Reva CLI with prompts
 * Enh #1079: Get file info using fxids from EOS
 * Enh #1088: Update LDAP user driver
 * Enh #1114: System information metrics cleanup
 * Enh #1071: System information included in Prometheus metrics
 * Enh #1094: Add logic for resolving storage references over webdav

Details
-------

 * Bugfix #1099: Do not restore recycle entry on purge

   This PR fixes a bug in the EOSFS driver that was restoring a deleted entry when asking for its
   permanent purge. EOS does not have the functionality to purge individual files, but the whole
   recycle of the user.

   https://github.com/cs3org/reva/pull/1099

 * Bugfix #1091: Allow listing the trashbin

   The trashbin endpoint expects the userid, not the username.

   https://github.com/cs3org/reva/pull/1091

 * Bugfix #1103: Restore and delete trash items via ocs

   The OCS api was not passing the correct key and references to the CS3 API. Furthermore, the
   owncloud storage driver was constructing the wrong target path when restoring.

   https://github.com/cs3org/reva/pull/1103

 * Bugfix #1090: Ensure ignoring public stray shares

   When using the json public shares manager, it can be the case we found a share with a resource_id
   that no longer exists.

   https://github.com/cs3org/reva/pull/1090

 * Bugfix #1064: Ensure ignoring stray shares

   When using the json shares manager, it can be the case we found a share with a resource_id that no
   longer exists. This PR aims to address his case by changing the contract of getPath and return
   the result of the STAT call instead of a generic error, so we can check the cause.

   https://github.com/cs3org/reva/pull/1064

 * Bugfix #1082: Minor fixes in reva cmd, gateway uploads and smtpclient

   https://github.com/cs3org/reva/pull/1082
   https://github.com/cs3org/reva/pull/1116

 * Bugfix #1115: Owncloud driver - propagate mtime on RemoveGrant

   When removing a grant the mtime change also needs to be propagated. Only affectsn the owncluod
   storage driver.

   https://github.com/cs3org/reva/pull/1115

 * Bugfix #1111: Handle redirection prefixes when extracting destination from URL

   The move function handler in ocdav extracts the destination path from the URL by removing the
   base URL prefix from the URL path. This would fail in case there is a redirection prefix. This PR
   takes care of that and it also allows zero-size uploads for localfs.

   https://github.com/cs3org/reva/pull/1111

 * Enhancement #1101: Add UID and GID in ldap auth driver

   The PR https://github.com/cs3org/reva/pull/1088/ added the functionality to lookup UID
   and GID from the ldap user provider. This PR adds the same to the ldap auth manager.

   https://github.com/cs3org/reva/pull/1101

 * Enhancement #1077: Add calens check to verify changelog entries in CI

   https://github.com/cs3org/reva/pull/1077

 * Enhancement #1072: Refactor Reva CLI with prompts

   The current CLI is a bit cumbersome to use with users having to type commands containing all the
   associated config with no prompts or auto-completes. This PR refactors the CLI with these
   functionalities.

   https://github.com/cs3org/reva/pull/1072

 * Enhancement #1079: Get file info using fxids from EOS

   This PR supports getting file information from EOS using the fxid value.

   https://github.com/cs3org/reva/pull/1079

 * Enhancement #1088: Update LDAP user driver

   The LDAP user driver can now fetch users by a single claim / attribute. Use an `attributefilter`
   like `(&(objectclass=posixAccount)({{attr}}={{value}}))` in the driver section.

   It also adds the uid and gid to the users opaque properties so that eos can use them for chown and
   acl operations.

   https://github.com/cs3org/reva/pull/1088

 * Enhancement #1114: System information metrics cleanup

   The system information metrics are now based on OpenCensus instead of the Prometheus client
   library. Furthermore, its initialization was moved out of the Prometheus HTTP service to keep
   things clean.

   https://github.com/cs3org/reva/pull/1114

 * Enhancement #1071: System information included in Prometheus metrics

   System information is now included in the main Prometheus metrics exposed at `/metrics`.

   https://github.com/cs3org/reva/pull/1071

 * Enhancement #1094: Add logic for resolving storage references over webdav

   This PR adds the functionality to resolve webdav references using the ocs webdav service by
   passing the resource's owner's token. This would need to be changed in production.

   https://github.com/cs3org/reva/pull/1094


Changelog for reva 1.1.0 (2020-08-11)
=======================================

The following sections list the changes in reva 1.1.0 relevant to
reva users. The changes are ordered by importance.

Summary
-------

 * Fix #1069: Pass build time variables while compiling
 * Fix #1047: Fix missing idp check in GetUser of demo userprovider
 * Fix #1038: Do not stat shared resources when downloading
 * Fix #1034: Fixed some error reporting strings and corresponding logs
 * Fix #1046: Fixed resolution of fileid in GetPathByID
 * Fix #1052: Ocfs: Lookup user to render template properly
 * Fix #1024: Take care of trailing slashes in OCM package
 * Fix #1025: Use lower-case name for changelog directory
 * Fix #1042: List public shares only created by the current user
 * Fix #1051: Disallow sharing the shares directory
 * Enh #1035: Refactor AppProvider workflow
 * Enh #1059: Improve timestamp precision while logging
 * Enh #1037: System information HTTP service
 * Enh #995: Add UID and GID to the user object from user package

Details
-------

 * Bugfix #1069: Pass build time variables while compiling

   We provide the option of viewing various configuration and version options in both reva CLI as
   well as the reva daemon, but we didn't actually have these values in the first place. This PR adds
   that info at compile time.

   https://github.com/cs3org/reva/pull/1069

 * Bugfix #1047: Fix missing idp check in GetUser of demo userprovider

   We've added a check for matching idp in the GetUser function of the demo userprovider

   https://github.com/cs3org/reva/issues/1047

 * Bugfix #1038: Do not stat shared resources when downloading

   Previously, we statted the resources in all download requests resulting in failures when
   downloading references. This PR fixes that by statting only in case the resource is not present
   in the shares folder. It also fixes a bug where we allowed uploading to the mount path, resulting
   in overwriting the user home directory.

   https://github.com/cs3org/reva/pull/1038

 * Bugfix #1034: Fixed some error reporting strings and corresponding logs

   https://github.com/cs3org/reva/pull/1034

 * Bugfix #1046: Fixed resolution of fileid in GetPathByID

   Following refactoring of fileid generations in the local storage provider, this ensures
   fileid to path resolution works again.

   https://github.com/cs3org/reva/pull/1046

 * Bugfix #1052: Ocfs: Lookup user to render template properly

   Currently, the username is used to construct paths, which breaks when mounting the `owncloud`
   storage driver at `/oc` and then expecting paths that use the username like
   `/oc/einstein/foo` to work, because they will mismatch the path that is used from propagation
   which uses `/oc/u-u-i-d` as the root, giving an `internal path outside root` error

   https://github.com/cs3org/reva/pull/1052

 * Bugfix #1024: Take care of trailing slashes in OCM package

   Previously, we assumed that the OCM endpoints would have trailing slashes, failing in case
   they didn't. This PR fixes that.

   https://github.com/cs3org/reva/pull/1024

 * Bugfix #1025: Use lower-case name for changelog directory

   When preparing a new release, the changelog entries need to be copied to the changelog folder
   under docs. In a previous change, all these folders were made to have lower case names,
   resulting in creation of a separate folder.

   https://github.com/cs3org/reva/pull/1025

 * Bugfix #1042: List public shares only created by the current user

   When running ocis, the public links created by a user are visible to all the users under the
   'Shared with others' tab. This PR fixes that by returning only those links which are created by a
   user themselves.

   https://github.com/cs3org/reva/pull/1042

 * Bugfix #1051: Disallow sharing the shares directory

   Previously, it was possible to create public links for and share the shares directory itself.
   However, when the recipient tried to accept the share, it failed. This PR prevents the creation
   of such shares in the first place.

   https://github.com/cs3org/reva/pull/1051

 * Enhancement #1035: Refactor AppProvider workflow

   Simplified the app-provider configuration: storageID is worked out automatically and UIURL
   is suppressed for now. Implemented the new gRPC protocol from the gateway to the appprovider.

   https://github.com/cs3org/reva/pull/1035

 * Enhancement #1059: Improve timestamp precision while logging

   Previously, the timestamp associated with a log just had the hour and minute, which made
   debugging quite difficult. This PR increases the precision of the associated timestamp.

   https://github.com/cs3org/reva/pull/1059

 * Enhancement #1037: System information HTTP service

   This service exposes system information via an HTTP endpoint. This currently only includes
   Reva version information but can be extended easily. The information are exposed in the form of
   Prometheus metrics so that we can gather these in a streamlined way.

   https://github.com/cs3org/reva/pull/1037

 * Enhancement #995: Add UID and GID to the user object from user package

   Currently, the UID and GID for users need to be read from the local system which requires local
   users to be present. This change retrieves that information from the user and auth packages and
   adds methods to retrieve it.

   https://github.com/cs3org/reva/pull/995


Changelog for reva 1.0.0 (2020-07-28)
=======================================

The following sections list the changes in reva 1.0.0 relevant to
reva users. The changes are ordered by importance.

Summary
-------

 * Fix #941: Fix initialization of json share manager
 * Fix #1006: Check if SMTP credentials are nil
 * Chg #965: Remove protocol from partner domains to match gocdb config
 * Enh #986: Added signing key capability
 * Enh #922: Add tutorial for deploying WOPI and Reva locally
 * Enh #979: Skip changelog enforcement for bot PRs
 * Enh #965: Enforce adding changelog in make and CI
 * Enh #1016: Do not enforce changelog on release
 * Enh #969: Allow requests to hosts with unverified certificates
 * Enh #914: Make httpclient configurable
 * Enh #972: Added a site locations exporter to Mentix
 * Enh #1000: Forward share invites to the provider selected in meshdirectory
 * Enh #1002: Pass the link to the meshdirectory service in token mail
 * Enh #1008: Use proper logging for ldap auth requests
 * Enh #970: Add required headers to SMTP client to prevent being tagged as spam
 * Enh #996: Split LDAP user filters
 * Enh #1007: Update go-tus version
 * Enh #1004: Update github.com/go-ldap/ldap to v3
 * Enh #974: Add functionality to create webdav references for OCM shares

Details
-------

 * Bugfix #941: Fix initialization of json share manager

   When an empty shares.json file existed the json share manager would fail while trying to
   unmarshal the empty file.

   https://github.com/cs3org/reva/issues/941
   https://github.com/cs3org/reva/pull/940

 * Bugfix #1006: Check if SMTP credentials are nil

   Check if SMTP credentials are nil before passing them to the SMTPClient, causing it to crash.

   https://github.com/cs3org/reva/pull/1006

 * Change #965: Remove protocol from partner domains to match gocdb config

   Minor changes for OCM cross-partner testing.

   https://github.com/cs3org/reva/pull/965

 * Enhancement #986: Added signing key capability

   The ocs capabilities can now hold the boolean flag to indicate url signing endpoint and
   middleware are available

   https://github.com/cs3org/reva/pull/986

 * Enhancement #922: Add tutorial for deploying WOPI and Reva locally

   Add a new tutorial on how to run Reva and Wopiserver together locally

   https://github.com/cs3org/reva/pull/922

 * Enhancement #979: Skip changelog enforcement for bot PRs

   Skip changelog enforcement for bot PRs.

   https://github.com/cs3org/reva/pull/979

 * Enhancement #965: Enforce adding changelog in make and CI

   When adding a feature or fixing a bug, a changelog needs to be specified, failing which the build
   wouldn't pass.

   https://github.com/cs3org/reva/pull/965

 * Enhancement #1016: Do not enforce changelog on release

   While releasing a new version of Reva, make release was failing because it was enforcing a
   changelog entry.

   https://github.com/cs3org/reva/pull/1016

 * Enhancement #969: Allow requests to hosts with unverified certificates

   Allow OCM to send requests to other mesh providers with the option of skipping certificate
   verification.

   https://github.com/cs3org/reva/pull/969

 * Enhancement #914: Make httpclient configurable

   - Introduce Options for the httpclient (#914)

   https://github.com/cs3org/reva/pull/914

 * Enhancement #972: Added a site locations exporter to Mentix

   Mentix now offers an endpoint that exposes location information of all sites in the mesh. This
   can be used in Grafana's world map view to show the exact location of every site.

   https://github.com/cs3org/reva/pull/972

 * Enhancement #1000: Forward share invites to the provider selected in meshdirectory

   Added a share invite forward OCM endpoint to the provider links (generated when a user picks a
   target provider in the meshdirectory service web interface), together with an invitation
   token and originating provider domain passed to the service via query params.

   https://github.com/sciencemesh/sciencemesh/issues/139
   https://github.com/cs3org/reva/pull/1000

 * Enhancement #1002: Pass the link to the meshdirectory service in token mail

   Currently, we just forward the token and the original user's domain when forwarding an OCM
   invite token and expect the user to frame the forward invite URL. This PR instead passes the link
   to the meshdirectory service, from where the user can pick the provider they want to accept the
   invite with.

   https://github.com/sciencemesh/sciencemesh/issues/139
   https://github.com/cs3org/reva/pull/1002

 * Enhancement #1008: Use proper logging for ldap auth requests

   Instead of logging to stdout we now log using debug level logging or error level logging in case
   the configured system user cannot bind to LDAP.

   https://github.com/cs3org/reva/pull/1008

 * Enhancement #970: Add required headers to SMTP client to prevent being tagged as spam

   Mails being sent through the client, specially through unauthenticated SMTP were being
   tagged as spam due to missing headers.

   https://github.com/cs3org/reva/pull/970

 * Enhancement #996: Split LDAP user filters

   The current LDAP user and auth filters only allow a single `%s` to be replaced with the relevant
   string. The current `userfilter` is used to lookup a single user, search for share recipients
   and for login. To make each use case more flexible we split this in three and introduced
   templates.

   For the `userfilter` we moved to filter templates that can use the CS3 user id properties
   `{{.OpaqueId}}` and `{{.Idp}}`: ``` userfilter =
   "(&(objectclass=posixAccount)(|(ownclouduuid={{.OpaqueId}})(cn={{.OpaqueId}})))"
   ```

   We introduced a new `findfilter` that is used when searching for users. Use it like this: ```
   findfilter =
   "(&(objectclass=posixAccount)(|(cn={{query}}*)(displayname={{query}}*)(mail={{query}}*)))"
   ```

   Furthermore, we also introduced a dedicated login filter for the LDAP auth manager: ```
   loginfilter = "(&(objectclass=posixAccount)(|(cn={{login}})(mail={{login}})))" ```

   These filter changes are backward compatible: `findfilter` and `loginfilter` will be
   derived from the `userfilter` by replacing `%s` with `{{query}}` and `{{login}}`
   respectively. The `userfilter` replaces `%s` with `{{.OpaqueId}}`

   Finally, we changed the default attribute for the immutable uid of a user to
   `ms-DS-ConsistencyGuid`. See
   https://docs.microsoft.com/en-us/azure/active-directory/hybrid/plan-connect-design-concepts
   for the background. You can fall back to `objectguid` or even `samaccountname` but you will run
   into trouble when user names change. You have been warned.

   https://github.com/cs3org/reva/pull/996

 * Enhancement #1007: Update go-tus version

   The lib now uses go mod which should help golang to sort out dependencies when running `go mod
   tidy`.

   https://github.com/cs3org/reva/pull/1007

 * Enhancement #1004: Update github.com/go-ldap/ldap to v3

   In the current version of the ldap lib attribute comparisons are case sensitive. With v3
   `GetEqualFoldAttributeValue` is introduced, which allows a case insensitive comparison.
   Which AFAICT is what the spec says: see
   https://github.com/go-ldap/ldap/issues/129#issuecomment-333744641

   https://github.com/cs3org/reva/pull/1004

 * Enhancement #974: Add functionality to create webdav references for OCM shares

   Webdav references will now be created in users' shares directory with the target set to the
   original resource's location in their mesh provider.

   https://github.com/cs3org/reva/pull/974


Changelog for reva 0.1.0 (2020-03-18)
=======================================

The following sections list the changes in reva 0.1.0 relevant to
reva users. The changes are ordered by importance.

Summary
-------

 * Enh #402: Build daily releases
 * Enh #416: Improve developer experience
 * Enh #468: remove vendor support
 * Enh #545: simplify configuration
 * Enh #561: improve the documentation
 * Enh #562: support home storages

Details
-------

 * Enhancement #402: Build daily releases

   Reva was not building releases of commits to the master branch. Thanks to @zazola.

   Commit-based released are generated every time a PR is merged into master. These releases are
   available at: https://reva-releases.web.cern.ch

   https://github.com/cs3org/reva/pull/402

 * Enhancement #416: Improve developer experience

   Reva provided the option to be run with a single configuration file by using the -c config flag.

   This PR adds the flag -dev-dir than can point to a directory containing multiple config files.
   The reva daemon will launch a new process per configuration file.

   Kudos to @refs.

   https://github.com/cs3org/reva/pull/416

 * Enhancement #468: remove vendor support

   Because @dependabot cannot update in a clean way the vendor dependecies Reva removed support
   for vendored dependencies inside the project.

   Dependencies will continue to be versioned but they will be downloaded when compiling the
   artefacts.

   https://github.com/cs3org/reva/pull/468
   https://github.com/cs3org/reva/pull/524

 * Enhancement #545: simplify configuration

   Reva configuration was difficul as many of the configuration parameters were not providing
   sane defaults. This PR and the related listed below simplify the configuration.

   https://github.com/cs3org/reva/pull/545
   https://github.com/cs3org/reva/pull/536
   https://github.com/cs3org/reva/pull/568

 * Enhancement #561: improve the documentation

   Documentation has been improved and can be consulted here: https://reva.link

   https://github.com/cs3org/reva/pull/561
   https://github.com/cs3org/reva/pull/545
   https://github.com/cs3org/reva/pull/568

 * Enhancement #562: support home storages

   Reva did not have any functionality to handle home storages. These PRs make that happen.

   https://github.com/cs3org/reva/pull/562
   https://github.com/cs3org/reva/pull/510
   https://github.com/cs3org/reva/pull/493
   https://github.com/cs3org/reva/pull/476
   https://github.com/cs3org/reva/pull/469
   https://github.com/cs3org/reva/pull/436
   https://github.com/cs3org/reva/pull/571


Changelog for reva 0.0.1 (2019-10-24)
=======================================

The following sections list the changes in reva 0.0.1 relevant to
reva users. The changes are ordered by importance.

Summary
-------

 * Enh #334: Create release procedure for Reva

Details
-------

 * Enhancement #334: Create release procedure for Reva

   Reva did not have any procedure to release versions. This PR brings a new tool to release Reva
   versions (tools/release) and prepares the necessary files for artefact distributed made
   from Drone into Github pages.

   https://github.com/cs3org/reva/pull/334


