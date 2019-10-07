package gdcli

import (
	"fmt"
)

func (req *zoneResponseBody) display() {
	fmt.Printf(" %s\n", req.Name)
	fmt.Printf("     %s %s %s\n", "|-", "id:", req.ID)
	fmt.Printf("     %s %s %s\n", "|-", "current_version_id:", req.CurrentVersionID)
	fmt.Printf("     %s %s %s\n", "|-", "created_at:", req.CurrentVersion.CreatedAt)
	fmt.Printf("     %s %s\n", "`-", "current_version")
	fmt.Printf("           %s %s %s\n", "|-", "last_modified_at:", req.CurrentVersion.LastModifiedAt)
	fmt.Printf("           %s %s %s\n", "|-", "name:", req.CurrentVersion.Name)
	fmt.Printf("           %s %s %s\n", "|-", "created_at:", req.CurrentVersion.CreatedAt)
	fmt.Printf("           %s %s %s\n", "|-", "id:", req.CurrentVersion.ID)
	fmt.Printf("           %s %s %t\n", "`-", "Editablet:", req.CurrentVersion.Editable)
}

func (req *recordResponseBody) display() {
	fmt.Printf("  %s\n", req.Name)
	fmt.Printf("     %s %s %s\n", "|-", "ID:", req.ID)
	if req.Type == "NS" {
		for _, rs := range req.Records {
			fmt.Printf("     %s %s %s\n", "|-", "Nsdname:", rs.Nsdname)
		}
	} else if req.Type == "TXT" {
		for _, rs := range req.Records {
			fmt.Printf("     %s %s %s\n", "|-", "Data:", rs.Data)
		}
	} else if req.Type == "MX" {
		for _, rs := range req.Records {
			fmt.Printf("     %s %s %d\n", "|-", "Priority:", rs.Prio)
			fmt.Printf("     %s %s %s\n", "|-", "Exchange:", rs.Exchange)
		}
	} else if req.Type == "CNAME" {
		for _, rs := range req.Records {
			fmt.Printf("     %s %s %s\n", "|-", "CNAME:", rs.Cname)
		}
	} else {
		for _, rs := range req.Records {
			fmt.Printf("     %s %s %s\n", "|-", "IPAddress:", rs.Address)
		}
	}
	fmt.Printf("     %s %s %s\n", "|-", "Type:", req.Type)
	fmt.Printf("     %s %s %d\n", "|-", "TTL:", req.TTL)
	fmt.Printf("     %s %s %t\n", "`-", "EnableAlias:", req.EnableAlias)
}
