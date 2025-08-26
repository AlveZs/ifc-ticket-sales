'use client';

import { selectTicketTypeAction } from "@/app/actions";
import { dollarStringFormatter } from "@/utils";


export type TicketTypeSelectProps = {
  defaultValue: 'full' | 'half';
  price: number;
}

export type TicketType = 'full' | 'half'

export function TicketTypeSelect(props: TicketTypeSelectProps) {
  const formattedFullPrice = dollarStringFormatter(props.price);
  const formattedHalfPrice = dollarStringFormatter(props.price / 2);

  return (
    <>
      <label htmlFor="ticket-type">Choose the ticket type</label>
      <select
        name="ticket-type"
        id="ticket-type"
        className="mt-2 rounded-lg bg-input px-4 py-[14px]"
        defaultValue={props.defaultValue}
        onChange={async (event) => {
          await selectTicketTypeAction(event.target.value as TicketType);
        }}
      >
        <option value="full">Full - {formattedFullPrice}</option>
        <option value="half">Half - {formattedHalfPrice}</option>
      </select>
    </>
  );
}
