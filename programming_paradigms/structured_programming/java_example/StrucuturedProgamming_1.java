package java_example;

import java.util.Scanner;

public class StrucuturedProgamming_1 {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);

        System.out.print("Masukkan angka pertama: ");

        int angka1 = scanner.nextInt();

        System.out.print("Masukkan angka kedua: ");

        int angka2 = scanner.nextInt();

        System.out.println("Hasil Penjumlahan: " + tambah(angka1, angka2));
        System.out.println("Hasil Pengurangan: " + kurang(angka1, angka2));
        System.out.println("Hasil Perkalian: " + kali(angka1, angka2));

        scanner.close();
    }

    public static int tambah(int a, int b) {
        return a + b;
    }

    public static int kurang(int a, int b) {
        return a - b;
    }

    public static int kali(int a, int b) {
        return a * b;
    }
}