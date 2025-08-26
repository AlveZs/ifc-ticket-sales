import { SpotSeat } from "@/app/components/SpotSeat";
import { Title } from "@/app/components/Title";
import { EventModel, SpotModel } from "@/app/models";
import Link from "next/link";
import { TicketType, TicketTypeSelect } from "./TicketTypeSelect";
import { cookies } from "next/headers";
import { dollarStringFormatter } from "@/utils";

export async function getSpots(eventId: string): Promise<{
  event: EventModel;
  spots: SpotModel[];
}> {
  const response = await fetch(
    `${process.env.API_URL}/events/${eventId}/spots`
  );

  return response.json();
}

export default async function SpotsLayoutPage({
  params,
}: {
  params: { eventId: string };
}) {
  const { eventId } = await params
  const { event, spots } = await getSpots(eventId);

  const rowLetters = spots.map((spot) => spot.name[0]);

  const uniqueRows = rowLetters.filter(
    (row, index) => rowLetters.indexOf(row) === index
  );

  const spotGroupedByRow = uniqueRows.map((row) => {
    return {
      row,
      spots: [
        ...spots
          .filter((spot) => spot.name[0] === row)
          .sort((a, b) => {
            const aNumber = parseInt(a.name.slice(1));
            const bNumber = parseInt(b.name.slice(1));

            if (aNumber < bNumber) {
              return -1;
            }

            if (aNumber > bNumber) {
              return 1;
            }

            return 0;
          }),
      ],
    };
  });

  const cookieStore = await cookies();
  const selectedSpots: string[] = JSON.parse(cookieStore.get("spots")?.value || "[]");
  let totalPrice = selectedSpots.length * event.price;
  const cookieTicketTypeValue = cookieStore.get("ticketType")?.value;
  const ticketType: TicketType = cookieTicketTypeValue ? cookieTicketTypeValue as TicketType : 'full';

  if (ticketType === 'half') {
    totalPrice /= 2;
  }

  const formattedTotalPrice = dollarStringFormatter(totalPrice);

  return (
    <main className="mt-10">
      <div className="flex w-[1176px] max-w-full flex-row flex-wrap justify-center gap-x-8 rounded-2xl bg-secondary p-4 md:justify-normal">
        <img src="/image.png" alt="" />
        <div className="flex max-w-full flex-col gap-y-6">
          <div className="flex flex-col gap-y-2">
            <p className="text-sm font-semibold uppercase text-subtitle">
              {new Date(event.date).toLocaleDateString("pt-BR", {
                weekday: "long",
                day: "2-digit",
                month: "2-digit",
                year: "numeric",
              })}
            </p>
            <p className="text-2xl font-semibold">{event.name}</p>
            <p className="font-normal">{event.location}</p>
          </div>
          <div className="flex h-[128px] flex-wrap justify-between gap-y-5 gap-x-3">
            <div className="flex flex-col gap-y-2">
              <p className="font-semibold">Organizer</p>
              <p className="text-sm font-normal">{event.organization}</p>
            </div>
            <div className="flex flex-col gap-y-2">
              <p className="font-semibold">Rating</p>
              <p className="text-sm font-normal">{event.rating}</p>
            </div>
          </div>
        </div>
      </div>
      <Title className="mt-10">Choose your spot</Title>
      <div className="mt-6 flex flex-wrap justify-between">
        <div className="mb-4 flex w-full max-w-[650px] flex-col gap-y-8 rounded-2xl bg-secondary p-6">
          <div className="rounded-2xl bg-bar py-4 text-center text-[20px] font-bold uppercase text-white">
            Stage
          </div>
          <div className="md:w-full md:justify-normal">
            {spotGroupedByRow.map((row) => {
              return (
                <div
                  key={row.row}
                  className="flex flex-row gap-3 items-center mb-3"
                >
                  <div className="w-4">{row.row}</div>
                  <div className="ml-2 flex flex-row">
                    {row.spots.map((spot) => {
                      return (
                        <SpotSeat
                          key={spot.name}
                          spotId={spot.name}
                          spotLabel={spot.name.slice(1)}
                          eventId={event.id}
                          selected={selectedSpots.includes(spot.name)}
                          disabled={false}
                        />
                      );
                    })}
                  </div>
                </div>
              );
            })}
          </div>
          <div className="flex w-full flex-row justify-around">
            <div className="flex flex-row items-center">
              <span className="mr-1 block h-4 w-4 rounded-full bg-[#00A96E]"></span>
              Available
            </div>
            <div className="flex flex-row items-center">
              <span className="mr-1 block h-4 w-4 rounded-full bg-[#A6ADBB]"></span>
              Occupied
            </div>
            <div className="flex flex-row items-center">
              <span className="mr-1 block h-4 w-4 rounded-full bg-(--color-blurple)"></span>
              Selected
            </div>
          </div>
        </div>
        <div className="flex w-full max-w-[478px] flex-col gap-y-6 rounded-2xl bg-secondary px-4 py-6">
          <h1 className="text-[20px] font-semibold">
            Check out the event prices
          </h1>
          <p>
            Full: {"$ 100.00"} <br />
            Half price: {`$ 50.00`}
          </p>
          <div className="flex flex-col">
            <TicketTypeSelect defaultValue={ticketType} price={event.price}/>
          </div>
          <div>Total: {formattedTotalPrice}</div>
          <Link
            href="/checkout"
            className="rounded-lg bg-btn-primary py-5 text-sm font-semibold uppercase text-btn-primary text-center hover:bg-white"
          >
            Go to checkout
          </Link>
        </div>
      </div>
    </main>
  );
}
