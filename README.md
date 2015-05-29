# Markdown Based Online Documenation Hosting

The [mdocs.org](https://mdocs.org) provides an online markdown hosting
for source text located in other public locations like Github,
Bitbucket, Gitlab etc.  User can create account in the site with a
subdomain to host the documenation.  To register to the site, user
should provide full name, email address, password and subdomain.  To
login to the site, user can use the email address and password.  The
subdomain and email will be linked, so using one email address only
one subdomain can be created.  OpenID connect based authentication can
be provided in the future.

User can add a profile description which will be displayed in the
subdomain.  When adding new documentation, user should provide the
link and the source text location and public listing status.  Links
without public listing can be access with direct URL.  Accessing
subdomain also display all the public listed documentation links.
Accessing non existing sudmain should rediect to main site.  User can
edit and delete documentation links.  There will be a link to refresh
the content at the bottom of site - without much distraction to the
actaul content.  The content should have print media CSS style.
