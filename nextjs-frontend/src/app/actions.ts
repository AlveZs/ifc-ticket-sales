"use server";

import { revalidateTag } from "next/cache";
import { cookies } from "next/headers";
import { redirect } from "next/navigation";

export async function selectSpotAction(eventId: string, spotName: string) {
  const cookieStore = await cookies();

  const spots = JSON.parse(cookieStore.get("spots")?.value || "[]");
  spots.push(spotName);
  const uniqueSpots = spots.filter(
    (spot: string, index: number) => spots.indexOf(spot) === index
  );

  cookieStore.set("spots", JSON.stringify(uniqueSpots));
  cookieStore.set("eventId", eventId);
}

export async function unselectSpotAction(spotName: string) {
  const cookieStore = await cookies();

  const spots = JSON.parse(cookieStore.get("spots")?.value || "[]");
  const newSpots = spots.filter((spot: string) => spot !== spotName);
  cookieStore.set("spots", JSON.stringify(newSpots));
}

export async function clearSpotsAction() {
  const cookieStore = await cookies();
  cookieStore.set("spots", "[]");
  cookieStore.set("eventId", "");
}

export async function selectTicketTypeAction(ticketType: "full" | "half") {
  const cookieStore = await cookies();
  cookieStore.set("ticketType", ticketType);
}

export async function checkoutAction({
  cardHash,
  email,
}: {
  cardHash: string;
  email: string;
}) {
  const cookieStore = await cookies();
  const eventId = cookieStore.get("eventId")?.value;
  const spots = JSON.parse(cookieStore.get("spots")?.value || "[]");
  const ticketType = cookieStore.get("ticketType")?.value || "full";

  const response = await fetch(`${process.env.API_URL}/checkout`, {
    method: "POST",
    body: JSON.stringify({
      event_id: eventId,
      card_hash: cardHash,
      ticket_type: ticketType,
      spots,
      email,
    }),
    headers: {
      "Content-Type": "application/json",
    },
  });

  if (!response.ok) {
    return {error: "Purchase error!"}
  }

  revalidateTag(`/events/${eventId}`);
  redirect(`/checkout/${eventId}/success`);
}
