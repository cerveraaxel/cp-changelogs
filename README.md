# cp-changelogs
View the cPanel changelogs for specific build versions from the command line

Simply specify the full cPanel version with or without the "parent" version build (11.102.0.4 or 102.0.4).

Example:

```
[root ~]# curl localhost:8080/11.102.0.4
{
  "Version": "102.0.4",
  "Details": {
    "ReleaseDate": "2022-02-16",
    "CaseList": [
      "Fixed case BWG-2907: Delete expired ActiveSync feedback notices.",
      "Fixed case COBRA-13685: Create replacement APIs for managing accounts’ contact emails.",
      "Fixed case COBRA-13720: Display left navigation for WHM Jupiter home page.",
      "Fixed case COBRA-13721: Update WHM and cPanel Icons to have Accent Colors.",
      "Fixed case CPANEL-39568: Return useful error message when scripts/addpop is not run as root.",
      "Fixed case CPANEL-39705: Fix bug that prevented installing plugins on Ubuntu.",
      "Fixed case CPANEL-39740: Update cpanel-php to include zip support.",
      "Fixed case CPANEL-39848: Tweak application search algorithm for more accurate and expected search results.",
      "Fixed case CPANEL-39861: Avoid using verify.cpanel.net for license state where possible.",
      "Fixed case CPANEL-39887: Make SOA edits not postfix the dname onto mname and rname.",
      "Fixed case CPANEL-39896: Adds ability to adjust maximum line length for SMTP transports in the Exim Basic Editor UI.",
      "Fixed case DUCK-6708: Add touch file to prevent WordPress Toolkit from installing during the initial cPanel installation.",
      "Fixed case HB-6352: Update cpanel-php74 to 7.4.27-1.cp11102.",
      "Fixed case PH-17696: Remove the Paper Lantern deprecation message from the cPanel interface.",
      "Fixed case PH-17697: Remove the Solutions page from the cPanel interface."
    ]
  }
}
```

If nothing is passed, a helper usage message is returned.

```
[root ~]# curl localhost:8080/

cPanel changelogs API:

Perform a query with the full cPanel version including major, minor and build versions
for which you would like to receive changelogs for.


Basic usage:

curl api.cpanel.axelcervera.com/11.102.0.5


Using the "parent" version value is optional:

curl api.cpanel.axelververa.com/102.0.5
```
