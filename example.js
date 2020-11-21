import url from 'k6/x/url';

export default function () {
  var parsed = url.parse("user.local:8000/path")
  console.log(`Scheme: ${parsed.scheme} `)
  console.log(`Host: ${parsed.host}`)
  console.log(`Path: ${parsed.path}`)

  var normalized = url.normalize("localhost:80///x///y/z/../././index.html?b=y&a=x#t=20")
  console.log(normalized)

  var split = url.splitHostPort("localhost:80")
  console.log(`Host: ${split.host}`)
  console.log(`Port: ${split.port}`)

  console.log(url.resolve("example.com"))
}