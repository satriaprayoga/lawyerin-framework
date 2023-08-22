CREATE OR REPLACE FUNCTION order_by_distance(latitude numeric, longitude numeric)
 RETURNS TABLE(firm_name varchar, address varchar, city varchar, province varchar, lat numeric, lng numeric, distance numeric)
 LANGUAGE sql
AS $$
SELECT firm_name, address, city, province, lat, lng, 6371 * acos(cos(radians(latitude)) * cos(radians(lat)) * cos(radians(lng) - radians(longitude)) + sin(radians(latitude)) * sin(radians(lat))) AS distance
FROM firm
ORDER BY distance ASC;
$$
