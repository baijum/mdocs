# Markdown Docs

The [mdocs.org](https://mdocs.org) is an online markdown hosting site.
The source text should be located in other public locations like
Github, Bitbucket, Gitlab etc.  User can create account in the site to
host documenations.  To register in the site, user should provide
username, full name, email address and password.  User will get a
profile page with a short bio and links to their documenations.  When
adding new documentation, user should provide documentation page
location and link to source text.  User can also choose whether to
link the documentation from the profile page or not.  However, all the
links will be publicly accessible.  User can edit and delete
documentation pages.  When the source text is no more availale, the
documentation link will return 404 Not Found status.

## Configuration

```bash
export MDOCS_POSTGRES_USER=baiju
export MDOCS_POSTGRES_PASSWORD=passwd
export MDOCS_POSTGRES_DBNAME=mdocs
export MDOCS_POSTGRES_SSLMODE=disable
export MDOCS_RANDOM_SALT=randomsalt
```

## Contribute!

This project is looking for contributions.  If you are interested to
contribute, fork the project and send pull requests.  If you have any
questions, send it to me: baiju.m.mail@gmail.com

-- [Baiju Muthukadan](http://muthukadan.net)
