<?php 

class Address {
  public string $city, $province, $country;
}

$address1 = new Address();
$address1->city = "Subang";
$address1->province = "Jawa Barat";
$address1->country = "Indonesia";

$address2 = $address1;
/**
 * - Di php, konsep nya pass by reference
 * - sudah otomatis, mengarah ke alamat memori yang sama
 * - ketika meng-set $address2->city = "Bandung"
 * - otomatis, $address1->city berisi bandung
 */
$address2->city = "Bandung";

print_r([$address1, $address2]);
/**
 * Result :
 * Array
(
    [0] => Address Object
        (
            [city] => Bandung
            [province] => Jawa Barat
            [country] => Indonesia
        )

    [1] => Address Object
        (
            [city] => Bandung
            [province] => Jawa Barat
            [country] => Indonesia
        )

)
 */