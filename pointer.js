class Address {
  city = "";
  province = "";
  country = "";
}

let address1 = new Address();
address1.city = "Subang";
address1.province = "Jawa Barat";
address1.country = "Indonesia";

let address2 = address1;
/**
 * Pada javascript, menggunakan konsep pass by reference
 */
address2.city = "Bandung";

console.log(address1);
console.log(address2);

/**
 * Value
 * Address {
  city: 'Bandung',
  province: 'Jawa Barat',
  country: 'Indonesia'
}
Address {
  city: 'Bandung',
  province: 'Jawa Barat',
  country: 'Indonesia'
}
 */
