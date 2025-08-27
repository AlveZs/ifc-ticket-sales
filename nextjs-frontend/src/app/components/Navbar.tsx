import Image from "next/image";
import Link from "next/link";
import { clearSpotsAction } from "../actions";

export function Navbar() {
  return (
    <div className="flex max-w-full items-center justify-items-stretch rounded-2xl bg-[#1D232A] px-6 py-2 shadow-nav">
      <div className="flex grow items-center justify-center">
        <Link href="/" onClick={clearSpotsAction}>
          <Image
            src="/icon.svg"
            alt="Icon DevTicket"
            width={136}
            height={48}
            className="max-h-[48px]"
          />
        </Link>
      </div>
      <Link href={"/checkout"} className="min-h-6 min-w-6 grow-0 items-center">
        <Image
          src="/cart-outline.svg"
          alt="Cart Icon"
          width={24}
          height={24}
        />
      </Link>
    </div>
  );
}
