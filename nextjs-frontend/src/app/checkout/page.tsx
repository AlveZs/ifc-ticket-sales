import { cookies } from "next/headers";
import { Title } from "../components/Title";
import { EventModel } from "../models";
import { redirect } from "next/navigation";
import { dollarStringFormatter } from "@/utils";
import { CheckoutForm } from "./CheckoutForm";

export async function getEvent(eventId: string): Promise<EventModel> {
  const response = await fetch(`${process.env.API_URL}/events/${eventId}`, {
    next: {
      tags: [`events/${eventId}`],
    },
  });

  return response.json();
}

export default async function CheckoutPage() {
  const cookieStore = await cookies();
  const eventId = cookieStore.get("eventId")?.value;
  if (!eventId) {
    return redirect("/");
  }
  const event = await getEvent(eventId);

  const selectedSpots: string[] = JSON.parse(
    cookieStore.get("spots")?.value || "[]"
  );
  const totalPrice = selectedSpots.length * event.price;
  const ticketType = cookieStore.get("ticketType")?.value;
  const formattedTotalPrice = dollarStringFormatter(
    ticketType === "half" ? totalPrice / 2 : totalPrice
  );

  return (
    <main className="mt-10 flex flex-wrap justify-center md:justify-between">
      <div className="mb-4 flex max-h-[250px] w-full max-w-[478px] flex-col gap-y-6 rounded-2xl bg-secondary p-4">
        <Title>Purchase Summary</Title>
        <p className="font-semibold">
          {event.name}
          <br />
          {event.location}
          <br />
          {new Date(event.date).toLocaleDateString("en-US", {
            weekday: "long",
            day: "2-digit",
            month: "2-digit",
            year: "numeric",
          })}
        </p>
        <p className="font-semibold text-white">{formattedTotalPrice}</p>
      </div>

      <div className="w-full max-w-[650px] rounded-2xl bg-secondary p-4">
        <Title>Payment details</Title>
        <CheckoutForm className="mt-6 flex flex-col gap-y-3">
          <div className="flex flex-col">
            <label htmlFor="email ">E-mail</label>
            <input
              type="email"
              name="email"
              id="person_email"
              className="mt-2 border-solid p-2 h-10 bg-input"
            />
          </div>
          <div className="flex flex-col">
            <label htmlFor="card_name">Name on card</label>
            <input
              type="text"
              name="card_name"
              id="person_name"
              className="mt-2 border-solid p-2 h-10 bg-input"
            />
          </div>
          <div className="flex flex-col">
            <label htmlFor="card_number">Card number</label>
            <input
              type="card_number"
              name="cc"
              id="person_cc_number"
              className="mt-2 border-solid p-2 h-10 bg-input"
            />
          </div>
          <div className="flex flex-wrap sm:justify-between">
            <div className="flex w-full flex-col md:w-auto">
              <label htmlFor="expire">Expiration date</label>
              <input
                type="text"
                name="expiration_date"
                id="person_cc_expiration_date"
                className="mt-2 border-solid p-2 h-10 bg-input"
              />
            </div>
            <div className="flex w-full flex-col md:w-auto">
              <label htmlFor="cvv">Card Verification Value (CVV)</label>
              <input
                type="card_number"
                name="cvv"
                id="person_cc_cvv"
                className="mt-2 border-solid p-2 h-10 bg-input"
              />
            </div>
          </div>
          <button className="rounded-lg bg-btn-primary py-4 px-4 text-sm font-semibold uppercase text-btn-primary">
            Finish payment
          </button>
        </CheckoutForm>
      </div>
    </main>
  );
}
