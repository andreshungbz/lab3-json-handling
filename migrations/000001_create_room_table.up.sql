CREATE TABLE IF NOT EXISTS room (
    id bigserial PRIMARY KEY,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    room_number integer NOT NULL,
    room_type text NOT NULL CHECK (room_type IN ('single', 'double', 'suite')),
    max_occupancy integer NOT NULL,
    has_balcony boolean NOT NULL,
    available boolean NOT NULL DEFAULT true
);