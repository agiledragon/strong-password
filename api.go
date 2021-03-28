package sp

func StrongPasswordChecker(s string) int {
	password := newPassword(s)

	len_less_than_atom_rule := atom(len_less_than_matcher, insert_action)
	len_more_than_atom_rule := atom(len_more_than_matcher, delete_action)
	types_atom_rule := atom(types_matcher, replace_action)
	cont_atom_rule := atom(cont_matcher, replace_action)

	len_rule := anyof(repeat(len_less_than_atom_rule), repeat(len_more_than_atom_rule))
	types_rule := repeat(types_atom_rule)
	cont_rule := repeat(cont_atom_rule)

	rule := allof(len_rule, types_rule, cont_rule)
	steps := rule(password)
	return steps
}
