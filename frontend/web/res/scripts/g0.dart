import 'dart:html';
import 'dart:async';
import 'dart:convert';
import 'package:g0/g0.dart';
import 'package:query_string/query_string.dart';
import 'package:dart_config/config.dart';
import 'package:dart_config/loaders/config_loader_httprequest.dart';

G0 go;

void main() {

  loadConfig('config.json').then((Map config) {
    String query = window.location.search;
    Map queryData = QueryString.parse(query);
    String offset = queryData['offset'];

    Element container = querySelector('#container');
    go = new G0(container, config, offset: offset);

    },
    onError: (error) => print(
      'config.json not found. rename config-sample.json to config.json'
    )
  );

}

Future<Map> loadConfig([String filename="config.yaml"]) {
  var config = new Config(filename,
      new ConfigHttpRequestLoader(),
      new JsonConfigParser());

  return config.readConfig();
}

/**
 * dirty fix for this issue:
 * https://github.com/chrisbu/dart_config/pull/10
 *
 * TODO: remove this when the issue is closed
 */
class JsonConfigParser implements ConfigParser {
  Future<Map> parse(String configText) {
    var completer = new Completer<Map>();

    var map = JSON.decode(configText);
    completer.complete(map);

    return completer.future;
  }
}
