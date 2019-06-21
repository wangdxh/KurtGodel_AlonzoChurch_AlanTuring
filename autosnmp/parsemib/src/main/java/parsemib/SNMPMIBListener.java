// Generated from ./SNMPMIB.g4 by ANTLR 4.7.1
package parsemib;
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link SNMPMIBParser}.
 */
public interface SNMPMIBListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link SNMPMIBParser#moduleDefinition}.
	 * @param ctx the parse tree
	 */
	void enterModuleDefinition(SNMPMIBParser.ModuleDefinitionContext ctx);
	/**
	 * Exit a parse tree produced by {@link SNMPMIBParser#moduleDefinition}.
	 * @param ctx the parse tree
	 */
	void exitModuleDefinition(SNMPMIBParser.ModuleDefinitionContext ctx);
	/**
	 * Enter a parse tree produced by {@link SNMPMIBParser#tagDefault}.
	 * @param ctx the parse tree
	 */
	void enterTagDefault(SNMPMIBParser.TagDefaultContext ctx);
	/**
	 * Exit a parse tree produced by {@link SNMPMIBParser#tagDefault}.
	 * @param ctx the parse tree
	 */
	void exitTagDefault(SNMPMIBParser.TagDefaultContext ctx);
	/**
	 * Enter a parse tree produced by {@link SNMPMIBParser#extensionDefault}.
	 * @param ctx the parse tree
	 */
	void enterExtensionDefault(SNMPMIBParser.ExtensionDefaultContext ctx);
	/**
	 * Exit a parse tree produced by {@link SNMPMIBParser#extensionDefault}.
	 * @param ctx the parse tree
	 */
	void exitExtensionDefault(SNMPMIBParser.ExtensionDefaultContext ctx);
	/**
	 * Enter a parse tree produced by {@link SNMPMIBParser#moduleBody}.
	 * @param ctx the parse tree
	 */
	void enterModuleBody(SNMPMIBParser.ModuleBodyContext ctx);
	/**
	 * Exit a parse tree produced by {@link SNMPMIBParser#moduleBody}.
	 * @param ctx the parse tree
	 */
	void exitModuleBody(SNMPMIBParser.ModuleBodyContext ctx);
	/**
	 * Enter a parse tree produced by {@link SNMPMIBParser#exports}.
	 * @param ctx the parse tree
	 */
	void enterExports(SNMPMIBParser.ExportsContext ctx);
	/**
	 * Exit a parse tree produced by {@link SNMPMIBParser#exports}.
	 * @param ctx the parse tree
	 */
	void exitExports(SNMPMIBParser.ExportsContext ctx);
	/**
	 * Enter a parse tree produced by {@link SNMPMIBParser#symbolsExported}.
	 * @param ctx the parse tree
	 */
	void enterSymbolsExported(SNMPMIBParser.SymbolsExportedContext ctx);
	/**
	 * Exit a parse tree produced by {@link SNMPMIBParser#symbolsExported}.
	 * @param ctx the parse tree
	 */
	void exitSymbolsExported(SNMPMIBParser.SymbolsExportedContext ctx);
	/**
	 * Enter a parse tree produced by {@link SNMPMIBParser#imports}.
	 * @param ctx the parse tree
	 */
	void enterImports(SNMPMIBParser.ImportsContext ctx);
	/**
	 * Exit a parse tree produced by {@link SNMPMIBParser#imports}.
	 * @param ctx the parse tree
	 */
	void exitImports(SNMPMIBParser.ImportsContext ctx);
	/**
	 * Enter a parse tree produced by {@link SNMPMIBParser#symbolsImported}.
	 * @param ctx the parse tree
	 */
	void enterSymbolsImported(SNMPMIBParser.SymbolsImportedContext ctx);
	/**
	 * Exit a parse tree produced by {@link SNMPMIBParser#symbolsImported}.
	 * @param ctx the parse tree
	 */
	void exitSymbolsImported(SNMPMIBParser.SymbolsImportedContext ctx);
	/**
	 * Enter a parse tree produced by {@link SNMPMIBParser#symbolsFromModuleList}.
	 * @param ctx the parse tree
	 */
	void enterSymbolsFromModuleList(SNMPMIBParser.SymbolsFromModuleListContext ctx);
	/**
	 * Exit a parse tree produced by {@link SNMPMIBParser#symbolsFromModuleList}.
	 * @param ctx the parse tree
	 */
	void exitSymbolsFromModuleList(SNMPMIBParser.SymbolsFromModuleListContext ctx);
	/**
	 * Enter a parse tree produced by {@link SNMPMIBParser#symbolsFromModule}.
	 * @param ctx the parse tree
	 */
	void enterSymbolsFromModule(SNMPMIBParser.SymbolsFromModuleContext ctx);
	/**
	 * Exit a parse tree produced by {@link SNMPMIBParser#symbolsFromModule}.
	 * @param ctx the parse tree
	 */
	void exitSymbolsFromModule(SNMPMIBParser.SymbolsFromModuleContext ctx);
	/**
	 * Enter a parse tree produced by {@link SNMPMIBParser#globalModuleReference}.
	 * @param ctx the parse tree
	 */
	void enterGlobalModuleReference(SNMPMIBParser.GlobalModuleReferenceContext ctx);
	/**
	 * Exit a parse tree produced by {@link SNMPMIBParser#globalModuleReference}.
	 * @param ctx the parse tree
	 */
	void exitGlobalModuleReference(SNMPMIBParser.GlobalModuleReferenceContext ctx);
	/**
	 * Enter a parse tree produced by {@link SNMPMIBParser#symbolList}.
	 * @param ctx the parse tree
	 */
	void enterSymbolList(SNMPMIBParser.SymbolListContext ctx);
	/**
	 * Exit a parse tree produced by {@link SNMPMIBParser#symbolList}.
	 * @param ctx the parse tree
	 */
	void exitSymbolList(SNMPMIBParser.SymbolListContext ctx);
	/**
	 * Enter a parse tree produced by {@link SNMPMIBParser#symbol}.
	 * @param ctx the parse tree
	 */
	void enterSymbol(SNMPMIBParser.SymbolContext ctx);
	/**
	 * Exit a parse tree produced by {@link SNMPMIBParser#symbol}.
	 * @param ctx the parse tree
	 */
	void exitSymbol(SNMPMIBParser.SymbolContext ctx);
	/**
	 * Enter a parse tree produced by {@link SNMPMIBParser#assignmentList}.
	 * @param ctx the parse tree
	 */
	void enterAssignmentList(SNMPMIBParser.AssignmentListContext ctx);
	/**
	 * Exit a parse tree produced by {@link SNMPMIBParser#assignmentList}.
	 * @param ctx the parse tree
	 */
	void exitAssignmentList(SNMPMIBParser.AssignmentListContext ctx);
	/**
	 * Enter a parse tree produced by {@link SNMPMIBParser#assignment}.
	 * @param ctx the parse tree
	 */
	void enterAssignment(SNMPMIBParser.AssignmentContext ctx);
	/**
	 * Exit a parse tree produced by {@link SNMPMIBParser#assignment}.
	 * @param ctx the parse tree
	 */
	void exitAssignment(SNMPMIBParser.AssignmentContext ctx);
	/**
	 * Enter a parse tree produced by {@link SNMPMIBParser#typeAssignment}.
	 * @param ctx the parse tree
	 */
	void enterTypeAssignment(SNMPMIBParser.TypeAssignmentContext ctx);
	/**
	 * Exit a parse tree produced by {@link SNMPMIBParser#typeAssignment}.
	 * @param ctx the parse tree
	 */
	void exitTypeAssignment(SNMPMIBParser.TypeAssignmentContext ctx);
	/**
	 * Enter a parse tree produced by {@link SNMPMIBParser#valueAssignment}.
	 * @param ctx the parse tree
	 */
	void enterValueAssignment(SNMPMIBParser.ValueAssignmentContext ctx);
	/**
	 * Exit a parse tree produced by {@link SNMPMIBParser#valueAssignment}.
	 * @param ctx the parse tree
	 */
	void exitValueAssignment(SNMPMIBParser.ValueAssignmentContext ctx);
	/**
	 * Enter a parse tree produced by {@link SNMPMIBParser#asnType}.
	 * @param ctx the parse tree
	 */
	void enterAsnType(SNMPMIBParser.AsnTypeContext ctx);
	/**
	 * Exit a parse tree produced by {@link SNMPMIBParser#asnType}.
	 * @param ctx the parse tree
	 */
	void exitAsnType(SNMPMIBParser.AsnTypeContext ctx);
	/**
	 * Enter a parse tree produced by {@link SNMPMIBParser#builtinType}.
	 * @param ctx the parse tree
	 */
	void enterBuiltinType(SNMPMIBParser.BuiltinTypeContext ctx);
	/**
	 * Exit a parse tree produced by {@link SNMPMIBParser#builtinType}.
	 * @param ctx the parse tree
	 */
	void exitBuiltinType(SNMPMIBParser.BuiltinTypeContext ctx);
	/**
	 * Enter a parse tree produced by {@link SNMPMIBParser#valuetype}.
	 * @param ctx the parse tree
	 */
	void enterValuetype(SNMPMIBParser.ValuetypeContext ctx);
	/**
	 * Exit a parse tree produced by {@link SNMPMIBParser#valuetype}.
	 * @param ctx the parse tree
	 */
	void exitValuetype(SNMPMIBParser.ValuetypeContext ctx);
	/**
	 * Enter a parse tree produced by {@link SNMPMIBParser#modulevalueType}.
	 * @param ctx the parse tree
	 */
	void enterModulevalueType(SNMPMIBParser.ModulevalueTypeContext ctx);
	/**
	 * Exit a parse tree produced by {@link SNMPMIBParser#modulevalueType}.
	 * @param ctx the parse tree
	 */
	void exitModulevalueType(SNMPMIBParser.ModulevalueTypeContext ctx);
	/**
	 * Enter a parse tree produced by {@link SNMPMIBParser#objectvalueType}.
	 * @param ctx the parse tree
	 */
	void enterObjectvalueType(SNMPMIBParser.ObjectvalueTypeContext ctx);
	/**
	 * Exit a parse tree produced by {@link SNMPMIBParser#objectvalueType}.
	 * @param ctx the parse tree
	 */
	void exitObjectvalueType(SNMPMIBParser.ObjectvalueTypeContext ctx);
	/**
	 * Enter a parse tree produced by {@link SNMPMIBParser#objectidentifiervaluetype}.
	 * @param ctx the parse tree
	 */
	void enterObjectidentifiervaluetype(SNMPMIBParser.ObjectidentifiervaluetypeContext ctx);
	/**
	 * Exit a parse tree produced by {@link SNMPMIBParser#objectidentifiervaluetype}.
	 * @param ctx the parse tree
	 */
	void exitObjectidentifiervaluetype(SNMPMIBParser.ObjectidentifiervaluetypeContext ctx);
	/**
	 * Enter a parse tree produced by {@link SNMPMIBParser#valuedesctype}.
	 * @param ctx the parse tree
	 */
	void enterValuedesctype(SNMPMIBParser.ValuedesctypeContext ctx);
	/**
	 * Exit a parse tree produced by {@link SNMPMIBParser#valuedesctype}.
	 * @param ctx the parse tree
	 */
	void exitValuedesctype(SNMPMIBParser.ValuedesctypeContext ctx);
	/**
	 * Enter a parse tree produced by {@link SNMPMIBParser#syntaxdesctype}.
	 * @param ctx the parse tree
	 */
	void enterSyntaxdesctype(SNMPMIBParser.SyntaxdesctypeContext ctx);
	/**
	 * Exit a parse tree produced by {@link SNMPMIBParser#syntaxdesctype}.
	 * @param ctx the parse tree
	 */
	void exitSyntaxdesctype(SNMPMIBParser.SyntaxdesctypeContext ctx);
	/**
	 * Enter a parse tree produced by {@link SNMPMIBParser#indexdesctype}.
	 * @param ctx the parse tree
	 */
	void enterIndexdesctype(SNMPMIBParser.IndexdesctypeContext ctx);
	/**
	 * Exit a parse tree produced by {@link SNMPMIBParser#indexdesctype}.
	 * @param ctx the parse tree
	 */
	void exitIndexdesctype(SNMPMIBParser.IndexdesctypeContext ctx);
	/**
	 * Enter a parse tree produced by {@link SNMPMIBParser#unknowndesctype}.
	 * @param ctx the parse tree
	 */
	void enterUnknowndesctype(SNMPMIBParser.UnknowndesctypeContext ctx);
	/**
	 * Exit a parse tree produced by {@link SNMPMIBParser#unknowndesctype}.
	 * @param ctx the parse tree
	 */
	void exitUnknowndesctype(SNMPMIBParser.UnknowndesctypeContext ctx);
	/**
	 * Enter a parse tree produced by {@link SNMPMIBParser#octetStringType}.
	 * @param ctx the parse tree
	 */
	void enterOctetStringType(SNMPMIBParser.OctetStringTypeContext ctx);
	/**
	 * Exit a parse tree produced by {@link SNMPMIBParser#octetStringType}.
	 * @param ctx the parse tree
	 */
	void exitOctetStringType(SNMPMIBParser.OctetStringTypeContext ctx);
	/**
	 * Enter a parse tree produced by {@link SNMPMIBParser#objectidentifiertype}.
	 * @param ctx the parse tree
	 */
	void enterObjectidentifiertype(SNMPMIBParser.ObjectidentifiertypeContext ctx);
	/**
	 * Exit a parse tree produced by {@link SNMPMIBParser#objectidentifiertype}.
	 * @param ctx the parse tree
	 */
	void exitObjectidentifiertype(SNMPMIBParser.ObjectidentifiertypeContext ctx);
	/**
	 * Enter a parse tree produced by {@link SNMPMIBParser#sequencetype}.
	 * @param ctx the parse tree
	 */
	void enterSequencetype(SNMPMIBParser.SequencetypeContext ctx);
	/**
	 * Exit a parse tree produced by {@link SNMPMIBParser#sequencetype}.
	 * @param ctx the parse tree
	 */
	void exitSequencetype(SNMPMIBParser.SequencetypeContext ctx);
	/**
	 * Enter a parse tree produced by {@link SNMPMIBParser#typedefine}.
	 * @param ctx the parse tree
	 */
	void enterTypedefine(SNMPMIBParser.TypedefineContext ctx);
	/**
	 * Exit a parse tree produced by {@link SNMPMIBParser#typedefine}.
	 * @param ctx the parse tree
	 */
	void exitTypedefine(SNMPMIBParser.TypedefineContext ctx);
	/**
	 * Enter a parse tree produced by {@link SNMPMIBParser#integerType}.
	 * @param ctx the parse tree
	 */
	void enterIntegerType(SNMPMIBParser.IntegerTypeContext ctx);
	/**
	 * Exit a parse tree produced by {@link SNMPMIBParser#integerType}.
	 * @param ctx the parse tree
	 */
	void exitIntegerType(SNMPMIBParser.IntegerTypeContext ctx);
	/**
	 * Enter a parse tree produced by {@link SNMPMIBParser#namedNumberList}.
	 * @param ctx the parse tree
	 */
	void enterNamedNumberList(SNMPMIBParser.NamedNumberListContext ctx);
	/**
	 * Exit a parse tree produced by {@link SNMPMIBParser#namedNumberList}.
	 * @param ctx the parse tree
	 */
	void exitNamedNumberList(SNMPMIBParser.NamedNumberListContext ctx);
	/**
	 * Enter a parse tree produced by {@link SNMPMIBParser#namedNumber}.
	 * @param ctx the parse tree
	 */
	void enterNamedNumber(SNMPMIBParser.NamedNumberContext ctx);
	/**
	 * Exit a parse tree produced by {@link SNMPMIBParser#namedNumber}.
	 * @param ctx the parse tree
	 */
	void exitNamedNumber(SNMPMIBParser.NamedNumberContext ctx);
	/**
	 * Enter a parse tree produced by {@link SNMPMIBParser#signedNumber}.
	 * @param ctx the parse tree
	 */
	void enterSignedNumber(SNMPMIBParser.SignedNumberContext ctx);
	/**
	 * Exit a parse tree produced by {@link SNMPMIBParser#signedNumber}.
	 * @param ctx the parse tree
	 */
	void exitSignedNumber(SNMPMIBParser.SignedNumberContext ctx);
	/**
	 * Enter a parse tree produced by {@link SNMPMIBParser#sizeConstraint}.
	 * @param ctx the parse tree
	 */
	void enterSizeConstraint(SNMPMIBParser.SizeConstraintContext ctx);
	/**
	 * Exit a parse tree produced by {@link SNMPMIBParser#sizeConstraint}.
	 * @param ctx the parse tree
	 */
	void exitSizeConstraint(SNMPMIBParser.SizeConstraintContext ctx);
	/**
	 * Enter a parse tree produced by {@link SNMPMIBParser#constraint}.
	 * @param ctx the parse tree
	 */
	void enterConstraint(SNMPMIBParser.ConstraintContext ctx);
	/**
	 * Exit a parse tree produced by {@link SNMPMIBParser#constraint}.
	 * @param ctx the parse tree
	 */
	void exitConstraint(SNMPMIBParser.ConstraintContext ctx);
	/**
	 * Enter a parse tree produced by {@link SNMPMIBParser#rangeconstraint}.
	 * @param ctx the parse tree
	 */
	void enterRangeconstraint(SNMPMIBParser.RangeconstraintContext ctx);
	/**
	 * Exit a parse tree produced by {@link SNMPMIBParser#rangeconstraint}.
	 * @param ctx the parse tree
	 */
	void exitRangeconstraint(SNMPMIBParser.RangeconstraintContext ctx);
}