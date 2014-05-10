part of G0;

class LiveApi implements Api {

  String _host;
  LiveApi(this._host);

  Future<Map> getImages({String offset: '0', int count: 20}){
    if(offset == null || offset == ''){
      offset = '0';
    }

    String url = '${_host}/${offset}/${count}';
    Future requestFuture = HttpRequest.getString(url)
                          .then(_onDataLoaded)
                          .catchError(_handleError);
    return requestFuture;
  }

  Map _onDataLoaded(String responseString){
    Map response;
    try {
      response = JSON.decode(responseString);
    } catch(exception, stackTrace) {
      print("can't decode received JSON");
      print(responseString);
      response = null;
    }
    return response;
  }

  void _handleError(var error){
    print("can't connect to API");
  }
}

