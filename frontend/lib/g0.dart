library G0;

import 'dart:html';
import 'dart:async';
import 'dart:math';
import 'dart:convert';
import 'package:intl/intl.dart';

part 'app/centered-float-list.dart';
part 'app/api.dart';
part 'app/api/fixture-api.dart';
part 'app/api/live-api.dart';
part 'app/image-list.dart';
part 'app/infinite-load.dart';
part 'app/detail.dart';

class G0 {

  static const String DATE_FORMAT = 'dd.MM.yyyyy, HH:mm';

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
    _loadImages(offset, imageList.perPage).then((_){
      if(offset != null){
        imageList.detail.showByOffset(offset);
      }
    });


    infiniteLoad.onFire.listen((_){
      _loadImages(imageList.currentOffset, imageList.perPage);
    });

  }

  /**
   * Loads [count] images older then [offset] async and shows loading spinner.
   * Displays images after [api] call is finished.
   * Initializes [centeredFloatList] on first call.
   */
  Future _loadImages(String offset, int count){
    Completer completer = new Completer();
    imageList.showLoading();
    Future<Map> future = api.getImages(offset: offset, count: count);
    future.then((result){
      if(result != null){
        imageList.showImages(result);
      }
    }).then((_){
         if(!centeredFloatList.isInitialized){
           centeredFloatList.init();
         }
         imageList.hideLoading();
         infiniteLoad.updateTargetHeight();
         infiniteLoad.activate();
         completer.complete();
      }
    );
    return completer.future;
  }
}
