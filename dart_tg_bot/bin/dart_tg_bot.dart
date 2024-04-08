import 'dart:async';

import 'package:teledart/model.dart';
import 'package:teledart/teledart.dart';
import 'package:teledart/telegram.dart';
import 'dart:io' as io;
import 'package:http/http.dart' as http;
import 'dart:core';
import 'dart:io' as io;
import 'dart:http';
import 'package:http/http.dart' as http;

const BOT_TOKEN = '7115239167:AAFoBXCx2d2azkjeHqc9X76nzoK4LnOJmtI';

Future<void> main() async {

  final username = (await Telegram(BOT_TOKEN).getMe()).username;
  var teledart = TeleDart(BOT_TOKEN, Event(username!));

  teledart.start();

  teledart.onCommand('start')
    .listen((message) => message.reply('напиши мне что нибудь и я отправлю картинку с этим текстом'));

  teledart.onMessage()
    .listen((message) => {
        http.get(Uri.parse("http://localhost:8000/api?text=" + message.text.toString())).then((response) {
          print(response.bodyBytes);
          io.File file = io.File(message.chat.id.toString()+".png");
          file.writeAsBytes(response.bodyBytes).then((res) => {
            message.replyPhoto(io.File(message.chat.id.toString()+'.png')).then((res) => {
              file.delete()
            })
          });
        })
      });
}
