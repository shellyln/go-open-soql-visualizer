package visualizer_test

import (
	"fmt"
	"testing"

	"github.com/shellyln/go-open-soql-parser/soql/parser"
	"github.com/shellyln/go-open-soql-visualizer/soql/visualizer"
)

func TestParse2(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name     string
		args     args
		want     interface{}
		wantErr  bool
		dbgBreak bool
	}{{
		name: "1",
		args: args{s: `
			SELECT
			    acc.Id xid
			  --, con.foo__r.xxx
			  , foo__r.bar__r.zzz
			  , foo__r.yyy
			  , con.Name xname
			  , con.acc.ddd xddd
			  , CONCAT(TRIM(acc.Name), '/', TRIM(con.Name), 123.45, 0xacc0) cname
			  , FLAT(acc.Name)
			  , (SELECT Id FROM con.Departments where uuu=con.Zzz and vvv=con.Id) qwerty
			  , (select Id from r3.lkjh where name='www')
			FROM
			    Contact con
			  , con.Account acc
			  , PPP.QQQ.RRR r3
			WHERE
			    not (Name like 'a%' or Name like 'b%')
				and
				acc.Name in ('a', 'b', 'c', null)
				and
				acc.Id in ('a', 'b', 'c', null)
				and
				r3.Name in (select x,Id,Name,(select w from ghjksfd) from Contact)
				and
				Name > 0001-01-02
				and
				(((Name > 0001-01-02T01:01:01.123456789Z)
				or
				Name = :param1))
				and
				con.Name = acc.Name
				and
				LEN(con.Name) > 0
				and
				foo__r.bar__r.zzz = 1
			ORDER BY
			    acc.Name desc nulls last
			  --, acc.Id desc nulls last
			  , xid
			  , con.Name
			OFFSET 1000 LIMIT 100
			FOR update viewstat, tracking
		`},
		want:    nil,
		wantErr: false,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.dbgBreak {
				t.Log("debug")
			}

			got, err := parser.Parse(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				return
			}

			s := visualizer.Visualize(got)
			fmt.Print(s)
		})
	}
}

func TestParse3(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name     string
		args     args
		want     interface{}
		wantErr  bool
		dbgBreak bool
	}{{
		name: "1",
		args: args{s: `
			SELECT
			    acc.Id xid
			  --, con.foo__r.xxx
			  , foo__r.bar__r.zzz
			  , foo__r.yyy
			  , con.Name xname
			  , con.acc.ddd xddd
			  , CONCAT(TRIM(acc.Name), '/', TRIM(con.Name), 123.45, 0xacc0) cname
			  , FLAT(acc.Name)
			FROM
			    Contact con
			  , con.Account acc
			  , PPP.QQQ.RRR r3
			WHERE
			    not (Name like 'a%' or Name like 'b%')
				and
				acc.Name in ('a', 'b', 'c', null)
				and
				acc.Id in ('a', 'b', 'c', null)
				and
				r3.Name in (select x,Id,Name from Contact)
				and
				Name > 0001-01-02
				and
				(((Name > 0001-01-02T01:01:01.123456789Z)
				or
				Name = :param1))
				and
				con.Name = acc.Name
				and
				LEN(con.Name) > 0
			GROUP BY
			    acc.Name
			  --, acc.Id
			  , xid
			  , con.Name
			  , foo__r.bar__r.zzz
			  , foo__r.yyy
			  , con.acc.ddd
			HAVING
			    LEN(MAX(con.Name)) > FOO(0)
				and
			    LEN(MAX(con.Id)) > 0
			ORDER BY
			    acc.Name desc nulls last
			  --, acc.Id desc nulls last
			  , xid
			  , con.Name
			OFFSET 1000 LIMIT 100
			FOR update viewstat, tracking
		`},
		want:    nil,
		wantErr: false,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.dbgBreak {
				t.Log("debug")
			}

			got, err := parser.Parse(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				return
			}

			s := visualizer.Visualize(got)
			fmt.Print(s)
		})
	}
}

func TestParse4(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name     string
		args     args
		want     interface{}
		wantErr  bool
		dbgBreak bool
	}{{
		name: "1",
		args: args{s: `
		SELECT
		    Account.Id
		  , Account.Name
		  , Account.Owner.Name
		  , Id
		  , Name
		  , (SELECT Id, Name
			 FROM Account.Cases
			 WHERE Id in (SELECT CaseId FROM LiveChatTranscript WHERE StartTime = TODAY))
		FROM
			Contact
		WHERE
			(Account.Name like 'foo%'
			or
			Account.Name like 'bar%') and Account.Owner.Name = 'aaa'
		`},
		want:    nil,
		wantErr: false,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.dbgBreak {
				t.Log("debug")
			}

			got, err := parser.Parse(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				return
			}

			s := visualizer.Visualize(got)
			fmt.Print(s)
		})
	}
}
