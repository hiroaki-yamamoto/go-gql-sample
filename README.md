# Go with graphql example

I wrote a simple graphql app. It's very simple; it has 2 public mutations, 1
public query, and 1 private query.

## Why I write this?

To memorize what I did.

## How to...

It's based on [GQLGen] and [Gorm]. Check there if you want to play this "sandbox".

## What I found
Go itself is very smart and fast language, and would be a choice to write a
backend of various services, but I found several things.

* ~~**No Admin Panel**: I'm mainly using Python and Django, and it has Admin Panel
    as you can see, and someone **provided** admin panel in Go.
    Yes, **it is past**. Now, the panel is no longer working...~~
    Finally, I decided to use [Prisma]
* **Needs to implement authentication subsystem manually**: Django has
    authentication subsystem, but I couldn't find good auth frameworks for go.
    AND **[I made it]**.
* ~~**Weak ORM**: [Gorm] is excellent ORM framework and I like it, however
    auto-migration feature is not so smart unlike Django. **I think this problem
    can be solved by writing migration generator.**~~
    I decided to use [Prisma]
* **Where is ResponseWriter!?**: Because there's no auth framework, I wanted to
    implement session with JWT and cookie, however, [GQLGen] provides only
    `context.Context` type in the resolver. Therefore, I couldn't set cookie and
    header. How can I **set** cookie in resolver??
    How can I **set** header in resolver??

## Conclusion
So, as conclusion, I think GQLGen is excellent for GraphQL API
(**not backend system**), and it needs an implementation to generate schema
migration automatically and authentication framework.

## Give me advice
If you have nice idea to deal above, please let me know thru issues list, and
of course, Pull Request is more appreciated.

## License
See: [LICENSE.md](LICENSE.md)

[GQLGen]: https://github.com/99designs/gqlgen
[Gorm]: https://github.com/jinzhu/gorm
[Prisma]: https://www.prisma.io/
[I made it]: https://github.com/hiroaki-yamamoto/gauth
