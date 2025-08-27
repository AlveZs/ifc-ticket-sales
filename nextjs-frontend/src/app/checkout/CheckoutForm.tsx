'use client';

import { PropsWithChildren } from "react";
import { checkoutAction } from "../actions";

export type CardInfos = {
  cardName: string;
  cardNumber: string;
  expireDate: string;
  cvv: string;
}

export async function getCardHash({ cardName, cardNumber, expireDate, cvv }: CardInfos) {
  const cardInfos: CardInfos = {
    cardName,
    cardNumber,
    expireDate,
    cvv
  };
  return Math.random().toString(36).substring(7);
}

export type CheckoutFormProps = {
  className?: string;
};

export function CheckoutForm(props: PropsWithChildren<CheckoutFormProps>) {
  return (
    <form action={async (formData: FormData) => {
      const card_hash = await getCardHash({
        cardName: formData.get("cardName") as string,
        cardNumber: formData.get("cardNumber") as string,
        expireDate: formData.get("expireDate") as string,
        cvv: formData.get("cvv") as string,
      });
      await checkoutAction({
        cardHash: card_hash,
        email: formData.get("email") as string
      });
    }}
    className={props.className}
    >
      <input type="hidden" name="card_hash" />
      {props.children}
    </form>
  );
}