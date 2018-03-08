var uri = context.getVariable("request.uri");
uri = uri.replace(context.getVariable("proxy.basepath"),"");
var path = /v4\/breakingnews\/(.*)(\?)(.*)/;
if (path.test(uri)) {
    uri = uri.replace(path, "v4/breakingnews/$1?$3&global=false");
}
context.setVariable("target.url", context.getVariable("target.url") + uri);
