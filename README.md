# antlion

Antlion is a WIP but usable email generator. Using a collection of "blocks" you can compose emails that will render reliably in most-all email clients, in plain-text and HTML.

Development of Antlion is largely handled by [Hussle Inc](github.com/hussleinc) employees.

## blocks

-   alert
-   button
-   code (like verify code, 123456 etc)
-   divider (hr)
-   header (h1)
-   markdown (paragraph + blackfriday parsing)
-   paragraph

## quick-start

```go
email := antlion.NewEmail(
	"Hello World!",
	"Preheader text",
	antlion.Header("Hello,"),
	antlion.Paragraph("Lorem ipsum...")
)

textContent := email.RenderText()
htmlContent := email.RenderHTML()

... whatever ...
```

## contributing

Feel free to contribute new components, please create a PR with screenshots. If you have access to something like [Litmus](litmus.com) or [Email on Acid](emailonacid.com), please use these tools first before submitting a PR, as we will _not_ merge content that does not acceptably render on all devices.
