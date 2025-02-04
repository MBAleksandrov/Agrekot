import 'package:http/http.dart' as http;
import 'dart:convert';


class ApiService {
  static const String _baseUrl = 'http://localhost:8080/auth/register';

  static Future<String> registerUser(String email, String password, String phone) async {
    final url = Uri.parse(_baseUrl);
    final headers = {'Content-Type': 'application/json'};
    final body = json.encode({
      'email': email,
      'password_hash': password,
      'phone': phone,
    });

    try {
      final response = await http.post(url, headers: headers, body: body);

      if (response.statusCode == 200) {
        return 'Пользователь успешно зарегистрирован';
      } else {
        return 'Ошибка регистрации. Статус: ${response.statusCode}';
      }
    } catch (e) {
      return 'Ошибка: $e';
    }
  }
}
