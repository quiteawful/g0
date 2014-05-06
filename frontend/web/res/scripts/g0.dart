import 'dart:html';
import 'package:g0/g0.dart';
import 'package:query_string/query_string.dart';

G0 go;

void main() {
  Api api = new FixtureApi();

  String query = window.location.search;
  Map queryData = QueryString.parse(query);
  String offset = queryData['offset'];

  Element container = querySelector('#container');
  go = new G0(container, api, offset: offset);
}
