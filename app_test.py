from app import APP
import unittest


class FlaskBookshelfTests(unittest.TestCase):

    @classmethod
    def setUpClass(cls):
        pass

    @classmethod
    def tearDownClass(cls):
        pass

    def setUp(self):
        self.APP = APP.test_client()
        self.APP.testing = True

    def tearDown(self):
        pass

    def test_home_status_code(self):
        result = self.APP.get('/')
        self.assertEqual(result.status_code, 200)

    def test_home_data(self):
        result = self.APP.get('/')
        self.assertEqual(result.data, b"Hello!")

    def test_blue_status_code(self):
        result = self.APP.get('/blue')
        self.assertEqual(result.status_code, 200)

    def test_blue_data(self):
        result = self.APP.get('/blue')
        self.assertEqual(result.data, b"Blue")

    def test_green_status_code(self):
        result = self.APP.get('/green')
        self.assertEqual(result.status_code, 200)

    def test_green_data(self):
        result = self.APP.get('/green')
        self.assertEqual(result.data, b"Green for dayz")


if __name__ == '__main__':
    unittest.main()
