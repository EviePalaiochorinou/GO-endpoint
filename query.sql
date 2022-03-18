/* 
1. Change the website field, so it only contains the domain
2. Count how many spots contain the same domain
3. Return spots which have a domain with a count greater than one */

select count(*), domains.domain from (
	SELECT split_part(domain, '/', 1) as domain from (
		select regexp_replace(website, '^(https?://)?(www\.)?','') as domain from "MY_TABLE" WHERE website IS NOT NULL AND website <> ''
		) as pre 
	) as domains
group by domains.domain having count(*) > 1
ORDER BY count DESC