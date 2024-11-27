import 'dart:convert';

import 'package:http/http.dart' as http;
import 'package:flutter_bloc/flutter_bloc.dart';

part 'event.dart';
part 'state.dart';

class CounterBloc extends Bloc<CounterEvent, CounterState> {
  CounterBloc() : super(CounterState(0)) {
    on<IncrementCount>(_incrementCount);
  }

  void _incrementCount(IncrementCount event, Emitter<CounterState> emit) async {
    try {
      final response =
          await http.get(Uri.parse("${Uri.base.origin}/api/increment"));
      emit(CounterState(jsonDecode(response.body)));
    } catch (e) {
      print(e);
    }
  }
}
