import 'package:flutter/material.dart';

void main() => runApp(MyApp());

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Flutter Registration Form',
      theme: ThemeData(
        primarySwatch: Colors.blue,
      ),
      home: RegistrationForm(),
    );
  }
}

class RegistrationForm extends StatefulWidget {
  @override
  _RegistrationFormState createState() => _RegistrationFormState();
}

class _RegistrationFormState extends State<RegistrationForm> {
  final _formKey = GlobalKey<FormState>();
  String _email = '';
  String _password = '';
  String _phone = '';

  void _submitForm() {
    if (_formKey.currentState?.validate() ?? false) {
      _formKey.currentState?.save();
      // Здесь можно отправить данные на сервер или выполнить другие действия
      print('Email: $_email');
      print('Password: $_password');
      print('Phone: $_phone');
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('Registration Form'),
      ),
      body: Padding(
        padding: const EdgeInsets.all(16.0),
        child: Form(
          key: _formKey,
          child: Column(
            children: [
              TextFormField(
                decoration: InputDecoration(labelText: 'Email'),
                keyboardType: TextInputType.emailAddress,
                validator: (value) {
                  if (value == null || value.isEmpty) {
                    return 'Please enter an email';
                  }
                  if (!RegExp(r'\S+@\S+\.\S+').hasMatch(value)) {
                    return 'Please enter a valid email address';
                  }
                  return null;
                },
                onSaved: (value) {
                  _email = value ?? '';
                },
              ),
              TextFormField(
                decoration: InputDecoration(labelText: 'Password'),
                obscureText: true,
                validator: (value) {
                  if (value == null || value.isEmpty) {
                    return 'Please enter a password';
                  }
                  return null;
                },
                onSaved: (value) {
                  _password = value ?? '';
                },
              ),
              TextFormField(
                decoration: InputDecoration(labelText: 'Phone'),
                obscureText: true,
                validator: (value) {
                  if (value == null || value.isEmpty) {
                    return 'Please enter your Phone';
                  }
                  return null;
                },
                onSaved: (value) {
                  _phone = value ?? '';
                },
              ),
              SizedBox(height: 20),
              ElevatedButton(
                onPressed: _submitForm,
                child: Text('Register'),
              ),
            ],
          ),
        ),
      ),
    );
  }
}