from app import app
import unittest

class FlaskBookshelfTests(unittest.TestCase):

    @classmethod
    def setUpClass(cls):
        pass

    @classmethod
    def tearDownClass(cls):
        pass

    def setUp(self):
        self.app = app.test_client()
        self.app.testing = True

    def tearDown(self):
        pass

    def test_home_status_code(self):
        result = self.app.get('/')
        self.assertEqual(result.status_code, 200)

    def test_home_data(self):
        result = self.app.get('/')
        self.assertEqual(result.data, b"Hello!")

    def test_blue_status_code(self):
        result = self.app.get('/blue')
        self.assertEqual(result.status_code, 200)

    def test_blue_data(self):
        result = self.app.get('/blue')
        self.assertEqual(result.data, b"Blue")

    def test_green_status_code(self):
        result = self.app.get('/green')
        self.assertEqual(result.status_code, 200)

    def test_green_data(self):
        result = self.app.get('/green')
        self.assertEqual(result.data, b"Green")

if __name__ == '__main__':
  unittest.main()
