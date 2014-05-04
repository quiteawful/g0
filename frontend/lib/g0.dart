library G0;

import 'dart:html';
import 'dart:async';

part 'app/centered-float-list.dart';
part 'app/api.dart';
part 'app/api/fixture-api.dart';
part 'app/api/live-api.dart';
part 'app/image-list.dart';
part 'app/infinite-load.dart';

class G0 {

  Element container;

  Api api;
  CenteredFloatList centeredFloatList;
  ImageList imageList;
  InfinteLoad infiniteLoad;

  /**
   * Initializes [G0] on [container] and loads first page from [api]
   * Named parameter [offset] is used for direkt linking
   */
  G0(this.container, this.api, {offset: null}){
    if(container == null){
      return;
    }
    Element imageListElement = container.querySelector('.image-list');
    imageList = new ImageList(imageListElement, 150, 150);
    centeredFloatList = new CenteredFloatList(imageListElement);
    infiniteLoad = new InfinteLoad(imageListElement);
    _loadImages(offset, imageList.perPage);

    infiniteLoad.onFire.listen((_){
      _loadImages(imageList.lastId, imageList.perPage);
    });
  }

  /**
   * Loads [count] images older then [offset] async and shows loading spinner.
   * Displays images after [api] call is finished.
   * Initializes [centeredFloatList] on first call.
   */
  void _loadImages(String offset, int count){
    imageList.showLoading();
    Future<Map> future = api.getImages(offset: offset, count: count);
    future.then((result) => imageList.showImages(result))
          .then((_){
             if(!centeredFloatList.isInitialized){
               centeredFloatList.init();
             }
             imageList.hideLoading();
             infiniteLoad.updateTargetHeight();
             infiniteLoad.activate();
          }
    );
  }
}
