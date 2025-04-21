WITH reference_vector AS (
  SELECT vec AS ref_vec
  FROM {{.Table}}
  LIMIT 1
)
SELECT /*+ recompile */ count(*)
FROM (
  SELECT /*+ no_merge */
    COSINE_DISTANCE(r.ref_vec, t.vec) as distance
  FROM {{.Table}} t, reference_vector r
) as results;

