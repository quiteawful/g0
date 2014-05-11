part of G0;

class LiveApi implements Api {

  String _host;
  int _reloadDelay = 0;
  int _loads = 0;

  Stopwatch _stopwatch = new Stopwatch()..start();

  LiveApi(this._host, this._reloadDelay){
    if(_reloadDelay == null){
      _reloadDelay = 0;
    }
  }

  Future<Map> getImages({String offset: '0', int count: 20}){
    if(_loads != 0 && _stopwatch.elapsedMilliseconds < _reloadDelay){
      return new Future((){ return null; });
    }
    _stopwatch.reset();
    _loads++;

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

