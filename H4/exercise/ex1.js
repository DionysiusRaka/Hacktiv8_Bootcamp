let varLet;
const varConst = "Constant variable";
console.log(varLet + " " + varConst);

let buah = ['apel', 'nanas', 'melon', 'semangka', 'salak'];
console.log(buah + "\n");

buah.push("Buah Naga");
console.log(buah + "\n");
buah.pop();
console.log(buah + "\n");

const saya = {
    nama_depan: "Dionysius",
    nama_belakang: "Raka",
    hobi: ["Ngopi", "Gaming", "Traveling"],
    angka_favorit: 9,
    memakai_kacamata: true
};
console.log("Nama saya " + saya.nama_depan + " " + saya.nama_belakang);

saya.angka_favorit = 8;
console.log(saya.angka_favorit);

saya.lulusan = "Hacktiv8";
console.log(saya);

Object.entries(saya.hobi).forEach(([key, value]) => {
    console.log(value);
});

Object.entries(saya).forEach(([key, value]) => {
    console.log(key);
});

Object.entries(saya).forEach(([key, value]) => {
    console.log(value);
});

cetakTanggal();

function cetakTanggal() {
    const tanggal = new Date().toDateString();
    console.log(tanggal);
}

ganjilGenap(2);
ganjilGenap(23);
ganjilGenap(20);
ganjilGenap(21);
ganjilGenap("21");

function ganjilGenap(inputA) {
    if (typeof inputA == 'number') {
        if (inputA % 2) console.log("Genap\n");
        else console.log("Ganjil");
    }
    else console.log("Invalid Data");
}
