import java.util.Scanner;

public class luasbidangdatar {
    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);
        int pilihan;

        while (true) {
            System.out.println("Program Hitung Luas:");
            System.out.println("1. Luas Persegi");
            System.out.println("2. Luas Persegi Panjang");
            System.out.println("3. Luas Segitiga");
            System.out.println("4. Luas Jajaran Genjang");
            System.out.println("5. Luas Trapesium");
            System.out.println("6. Luas Lingkaran");
            System.out.println("7. Luas Belah Ketupat");
            System.out.println("8. Luas Layang-layang");

            System.out.print("Masukkan Pilihan: ");
            pilihan = input.nextInt();

            if (pilihan == 1) {
                // Luas Persegi
                System.out.print("Masukkan Panjang Sisi: ");
                int sisi = input.nextInt();
                int luas = sisi * sisi;
                System.out.println("Luas Persegi: " + luas);
            } else if (pilihan == 2) {
                // Luas Persegi Panjang
                System.out.print("Masukkan Panjang: ");
                int panjang = input.nextInt();
                System.out.print("Masukkan Lebar: ");
                int lebar = input.nextInt();
                int luas = panjang * lebar;
                System.out.println("Luas Persegi Panjang: " + luas);
            } else if (pilihan == 3) {
                // Luas Segitiga
                System.out.print("Masukkan Alas: ");
                int alas = input.nextInt();
                System.out.print("Masukkan Tinggi: ");
                int tinggi = input.nextInt();
                int luas = (alas * tinggi) / 2;
                System.out.println("Luas Segitiga: " + luas);
            } else if (pilihan == 4) {
                // Luas Jajaran Genjang
                System.out.print("Masukkan Alas: ");
                int alas = input.nextInt();
                System.out.print("Masukkan Tinggi: ");
                int tinggi = input.nextInt();
                int luas = alas * tinggi;
                System.out.println("Luas Jajaran Genjang: " + luas);
            } else if (pilihan == 5) {
                // Luas Trapesium
                System.out.print("Masukkan Panjang Sisi Atas: ");
                int sisiAtas = input.nextInt();
                System.out.print("Masukkan Panjang Sisi Bawah: ");
                int sisiBawah = input.nextInt();
                System.out.print("Masukkan Tinggi: ");
                int tinggi = input.nextInt();
                int luas = ((sisiAtas + sisiBawah) * tinggi) / 2;
                System.out.println("Luas Trapesium: " + luas);
            } else if (pilihan == 6) {
                // Luas Lingkaran
                System.out.print("Masukkan Jari-jari: ");
                double jariJari = input.nextDouble();
                double luas = Math.PI * jariJari * jariJari;
                System.out.println("Luas Lingkaran: " + luas);
            } else if (pilihan == 7) {
                // Luas Belah Ketupat
                System.out.print("Masukkan Diagonal 1: ");
                double diagonal1 = input.nextDouble();
                System.out.print("Masukkan Diagonal 2: ");
                double diagonal2 = input.nextDouble();
                double luas = (diagonal1 * diagonal2) / 2;
                System.out.println("Luas Belah Ketupat: " + luas);
            } else if (pilihan == 8) {
                // Luas Layang-layang
                System.out.print("Masukkan Diagonal 1: ");
                double diagonal1 = input.nextDouble();
                System.out.print("Masukkan Diagonal 2: ");
                double diagonal2 = input.nextDouble();
                double luas = (diagonal1 * diagonal2) / 2;
                System.out.println("Luas Layang-layang: " + luas);

            } else {
                System.out.println("Pilihan tidak valid.");
            }

            System.out.println();

            System.out.print("Apakah Anda ingin menghitung lagi? (y/n): ");
            String ulang = input.next();
            if (ulang.equalsIgnoreCase("n")) {
                break;
            }
        }

        System.out.println("Program berakhir.");
    }
    
}
