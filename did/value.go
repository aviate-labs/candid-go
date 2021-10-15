package did

import (
	"math/big"
	"strconv"
	"strings"

	"github.com/aviate-labs/candid-go/idl"
	"github.com/aviate-labs/candid-go/internal/candidvalue"
	"github.com/di-wu/parser/ast"
)

func ConvertValues(n *ast.Node) ([]idl.Type, []interface{}, error) {
	switch n.Type {
	case candidvalue.BoolValueT:
		switch n.Value {
		case "true":
			return []idl.Type{new(idl.Bool)}, []interface{}{true}, nil
		case "false":
			return []idl.Type{new(idl.Bool)}, []interface{}{false}, nil
		default:
			panic(n)
		}
	case candidvalue.NullT:
		return []idl.Type{new(idl.Null)}, []interface{}{nil}, nil
	case candidvalue.NumT:
		typ, arg, err := convertNum(n)
		if err != nil {
			return nil, nil, err
		}
		return []idl.Type{typ}, []interface{}{arg}, nil
	case candidvalue.OptValueT:
		types, args, err := ConvertValues(n.Children()[0])
		if err != nil {
			return nil, nil, err
		}
		return []idl.Type{idl.NewOpt(types[0])}, []interface{}{args[0]}, nil
	case candidvalue.TextT:
		n := n.Children()[0]
		s := strings.TrimPrefix(strings.TrimSuffix(n.Value, "\""), "\"")
		return []idl.Type{new(idl.Text)}, []interface{}{s}, nil
	case candidvalue.ValuesT:
		var (
			types []idl.Type
			args  []interface{}
		)
		for _, n := range n.Children() {
			typ, arg, err := ConvertValues(n)
			if err != nil {
				return nil, nil, err
			}
			types = append(types, typ...)
			args = append(args, arg...)
		}
		return types, args, nil
	default:
		panic(n)
	}
}

func convertNum(n *ast.Node) (idl.Type, interface{}, error) {
	switch n := n.Children(); len(n) {
	case 1:
		n := n[0]

		// float64
		if strings.Contains(n.Value, ".") {
			v := strings.ReplaceAll(n.Value, "_", "")
			f, err := strconv.ParseFloat(v, 64)
			if err != nil {
				return nil, nil, err
			}
			return idl.Float64(), big.NewFloat(f), nil
		}

		// int
		v := strings.ReplaceAll(n.Value, "_", "")
		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return nil, nil, err
		}
		return new(idl.Int), big.NewInt(i), nil
	case 2:
		vArg := n[0].Value
		vType := n[1].Value

		// floats
		if vType == "float32" || vType == "float64" {
			v := strings.ReplaceAll(vArg, "_", "")
			f, err := strconv.ParseFloat(v, 64)
			if err != nil {
				return nil, nil, err
			}
			bf := big.NewFloat(f)
			switch n := n[1]; n.Value {
			case "float32":
				return idl.Float32(), bf, nil
			default:
				return idl.Float64(), bf, nil
			}
		}

		// ints
		v := strings.ReplaceAll(vArg, "_", "")
		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return nil, nil, err
		}
		bi := big.NewInt(i)
		switch vType {
		case "nat":
			return new(idl.Nat), bi, nil
		case "nat8":
			return idl.Nat8(), bi, nil
		case "nat16":
			return idl.Nat16(), bi, nil
		case "nat32":
			return idl.Nat32(), bi, nil
		case "nat64":
			return idl.Nat64(), bi, nil
		case "int":
			return new(idl.Int), bi, nil
		case "int8":
			return idl.Int8(), bi, nil
		case "int16":
			return idl.Int16(), bi, nil
		case "int32":
			return idl.Int32(), bi, nil
		case "int64":
			return idl.Int64(), bi, nil
		default:
			panic(n)
		}
	default:
		panic(n)
	}
}
