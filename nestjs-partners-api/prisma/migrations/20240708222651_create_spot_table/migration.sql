-- CreateEnum
CREATE TYPE "SpotStatus" AS ENUM ('available', 'reserved');

-- CreateTable
CREATE TABLE "Spot" (
    "id" TEXT NOT NULL,
    "name" TEXT NOT NULL,
    "status" "SpotStatus" NOT NULL,
    "eventId" TEXT NOT NULL,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) NOT NULL,

    CONSTRAINT "Spot_pkey" PRIMARY KEY ("id")
);

-- AddForeignKey
ALTER TABLE "Spot" ADD CONSTRAINT "Spot_eventId_fkey" FOREIGN KEY ("eventId") REFERENCES "Event"("id") ON DELETE RESTRICT ON UPDATE CASCADE;
