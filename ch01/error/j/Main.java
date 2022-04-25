package hoge;

import java.io.FileInputStream;
import java.io.IOException;

// file-open
public class Main {
    public static void main(String[] args) throws IOException {
        try (var in = new FileInputStream("hoge.txt")) {
            // inを使ったなにかの処理
        }
    }
}
// file-open
