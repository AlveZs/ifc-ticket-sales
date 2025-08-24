import { Title } from "../components/Title";
import { EventModel } from "../models";

export default function CheckoutPage() {
  const event: EventModel = {
    id: "1",
    name: "Test",
    organization: "Productions",
    date: "2022-12-31T00:00:00.000Z",
    location: "Salvador",
  };
  return (
    <main className="mt-10 flex flex-wrap justify-center md:justify-between">
      <div className="mb-4 flex max-h-[250px] w-full max-w-[478px] flex-col gap-y-6 rounded-2xl bg-secondary p-4">
        <Title>Purchase Summary</Title>
        <p className="font-semibold">
          {event.name}
          <br />
          {event.location}
          <br />
          {new Date(event.date).toLocaleDateString("pt-BR", {
            weekday: "long",
            day: "2-digit",
            month: "2-digit",
            year: "numeric",
          })}
        </p>
        <p className="font-semibold text-white">
          Total price
        </p>
      </div>
      <div className="w-full max-w-[650px] rounded-2xl bg-secondary p-4">
        <Title>Payment details</Title>

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
        <div className="flex flex-col">
          <label htmlFor="expire">Expiration date</label>
          <input
            type="text"
            name="expiration_date"
            id="person_cc_expiration_date"
            className="mt-2 border-solid p-2 h-10 bg-input"
          />
        </div>
        <div className="flex flex-col">
          <label htmlFor="cvv">Card Verification Value (CVV)</label>
          <input
            type="card_number"
            name="cvv"
            id="person_cc_cvv"
            className="mt-2 border-solid p-2 h-10 bg-input"
          />
        </div>
        <button className="rounded-lg bg-btn-primary py-4 px-4 text-sm font-semibold uppercase text-btn-primary">
          Finish payment
        </button>
      </div>
    </main>
  );
}
