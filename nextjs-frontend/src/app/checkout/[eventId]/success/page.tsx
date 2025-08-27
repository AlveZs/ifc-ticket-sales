import { clearSpotsAction } from "@/app/actions";
import { Title } from "@/app/components/Title";
import { EventModel } from "@/app/models";
import { localeDateFormatter } from "@/utils";
import { cookies } from "next/headers";
import Link from "next/link";

export async function getEvent(eventId: string): Promise<EventModel> {
  const response = await fetch(`${process.env.API_URL}/events/${eventId}`, {
    next: {
      tags: [`events/${eventId}`],
    },
  });

  return response.json();
}
export default async function CheckoutSuccessPage({
  params,
}: {
  params: { eventId: string };
}) {
  const { eventId } = await params;
  const event = await getEvent(eventId);
  const cookieStore = await cookies();
  const selectedSpots: string[] = JSON.parse(
    cookieStore.get("spots")?.value || "[]"
  );
  return (
    <main className="mt-10 flex flex-col flex-wrap items-center">
      <Title>Purchase completed successfully</Title>
      <div className="mb-4 flex max-h-[290px] w-full max-w-[478px] flex-col gap-y-6 rounded-2xl bg-secondary p-4">
        <Title>Purchase summary</Title>
        <p className="font-semibold">
          Event {event.name}
          <br />
          Location {event.location}
          <br />
          Date {localeDateFormatter(event.date)}
        </p>
        <p className="font-semibold text-white">
          Chosen spots: {selectedSpots.join(", ")}
        </p>
        <Link href={"/"} onClick={clearSpotsAction}>
          <button 
            className="rounded-lg w-full bg-btn-primary
              py-4 px-4 text-sm font-semibold uppercase text-btn-primary
              cursor-pointer
            ">
            OK
          </button>
        </Link>
      </div>
    </main>
  );
}
