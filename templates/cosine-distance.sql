
WITH reference_vector AS (
  SELECT vec AS ref_vec
  FROM {{.Table}}
  LIMIT 1
)
SELECT /*+ no_merge */
COSINE_DISTANCE(r.ref_vec, t.vec) as distance
FROM {{.Table}} t, reference_vector r

