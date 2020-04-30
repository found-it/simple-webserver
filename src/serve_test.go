package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestGreen(t *testing.T) {
    req, err := http.NewRequest("GET", "/green", nil)
    if err != nil {
        t.Fatal(err)
    }

    w := httptest.NewRecorder()
    handler := http.HandlerFunc(Green)

    handler.ServeHTTP(w, req)

    if status := w.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v | want %v", status, http.StatusOK)
    } else {
        t.Logf("handler returned correct status code: %v", status)
    }

    expected := `Green for dayz!`
    if w.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v | want %v", w.Body.String(), expected)
    } else {
        t.Logf("handler returned correct body: %v", w.Body.String())
    }
}


func TestBlue(t *testing.T) {
    req, err := http.NewRequest("GET", "/blue", nil)
    if err != nil {
        t.Fatal(err)
    }

    w := httptest.NewRecorder()
    handler := http.HandlerFunc(Blue)

    handler.ServeHTTP(w, req)

    if status := w.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v | want %v", status, http.StatusOK)
    } else {
        t.Logf("handler returned correct status code: %v", status)
    }

    expected := `Blue!`
    if w.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v | want %v", w.Body.String(), expected)
    } else {
        t.Logf("handler returned correct body: %v", w.Body.String())
    }
}


func TestGetIPs(t *testing.T) {
    req, err := http.NewRequest("GET", "/", nil)
    if err != nil {
        t.Fatal(err)
    }

    w := httptest.NewRecorder()
    handler := http.HandlerFunc(GetIPs)

    handler.ServeHTTP(w, req)

    if status := w.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v | want %v", status, http.StatusOK)
    } else {
        t.Logf("handler returned correct status code: %v", status)
    }
}
