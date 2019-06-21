// Generated from ./SNMPMIB.g4 by ANTLR 4.7.1
package parsemib;
import org.antlr.v4.runtime.tree.ParseTreeVisitor;

/**
 * This interface defines a complete generic visitor for a parse tree produced
 * by {@link SNMPMIBParser}.
 *
 * @param <T> The return type of the visit operation. Use {@link Void} for
 * operations with no return type.
 */
public interface SNMPMIBVisitor<T> extends ParseTreeVisitor<T> {
	/**
	 * Visit a parse tree produced by {@link SNMPMIBParser#moduleDefinition}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitModuleDefinition(SNMPMIBParser.ModuleDefinitionContext ctx);
	/**
	 * Visit a parse tree produced by {@link SNMPMIBParser#tagDefault}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitTagDefault(SNMPMIBParser.TagDefaultContext ctx);
	/**
	 * Visit a parse tree produced by {@link SNMPMIBParser#extensionDefault}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitExtensionDefault(SNMPMIBParser.ExtensionDefaultContext ctx);
	/**
	 * Visit a parse tree produced by {@link SNMPMIBParser#moduleBody}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitModuleBody(SNMPMIBParser.ModuleBodyContext ctx);
	/**
	 * Visit a parse tree produced by {@link SNMPMIBParser#exports}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitExports(SNMPMIBParser.ExportsContext ctx);
	/**
	 * Visit a parse tree produced by {@link SNMPMIBParser#symbolsExported}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitSymbolsExported(SNMPMIBParser.SymbolsExportedContext ctx);
	/**
	 * Visit a parse tree produced by {@link SNMPMIBParser#imports}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitImports(SNMPMIBParser.ImportsContext ctx);
	/**
	 * Visit a parse tree produced by {@link SNMPMIBParser#symbolsImported}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitSymbolsImported(SNMPMIBParser.SymbolsImportedContext ctx);
	/**
	 * Visit a parse tree produced by {@link SNMPMIBParser#symbolsFromModuleList}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitSymbolsFromModuleList(SNMPMIBParser.SymbolsFromModuleListContext ctx);
	/**
	 * Visit a parse tree produced by {@link SNMPMIBParser#symbolsFromModule}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitSymbolsFromModule(SNMPMIBParser.SymbolsFromModuleContext ctx);
	/**
	 * Visit a parse tree produced by {@link SNMPMIBParser#globalModuleReference}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitGlobalModuleReference(SNMPMIBParser.GlobalModuleReferenceContext ctx);
	/**
	 * Visit a parse tree produced by {@link SNMPMIBParser#symbolList}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitSymbolList(SNMPMIBParser.SymbolListContext ctx);
	/**
	 * Visit a parse tree produced by {@link SNMPMIBParser#symbol}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitSymbol(SNMPMIBParser.SymbolContext ctx);
	/**
	 * Visit a parse tree produced by {@link SNMPMIBParser#assignmentList}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitAssignmentList(SNMPMIBParser.AssignmentListContext ctx);
	/**
	 * Visit a parse tree produced by {@link SNMPMIBParser#assignment}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitAssignment(SNMPMIBParser.AssignmentContext ctx);
	/**
	 * Visit a parse tree produced by {@link SNMPMIBParser#typeAssignment}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitTypeAssignment(SNMPMIBParser.TypeAssignmentContext ctx);
	/**
	 * Visit a parse tree produced by {@link SNMPMIBParser#valueAssignment}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitValueAssignment(SNMPMIBParser.ValueAssignmentContext ctx);
	/**
	 * Visit a parse tree produced by {@link SNMPMIBParser#asnType}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitAsnType(SNMPMIBParser.AsnTypeContext ctx);
	/**
	 * Visit a parse tree produced by {@link SNMPMIBParser#builtinType}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitBuiltinType(SNMPMIBParser.BuiltinTypeContext ctx);
	/**
	 * Visit a parse tree produced by {@link SNMPMIBParser#valuetype}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitValuetype(SNMPMIBParser.ValuetypeContext ctx);
	/**
	 * Visit a parse tree produced by {@link SNMPMIBParser#modulevalueType}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitModulevalueType(SNMPMIBParser.ModulevalueTypeContext ctx);
	/**
	 * Visit a parse tree produced by {@link SNMPMIBParser#objectvalueType}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitObjectvalueType(SNMPMIBParser.ObjectvalueTypeContext ctx);
	/**
	 * Visit a parse tree produced by {@link SNMPMIBParser#objectidentifiervaluetype}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitObjectidentifiervaluetype(SNMPMIBParser.ObjectidentifiervaluetypeContext ctx);
	/**
	 * Visit a parse tree produced by {@link SNMPMIBParser#valuedesctype}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitValuedesctype(SNMPMIBParser.ValuedesctypeContext ctx);
	/**
	 * Visit a parse tree produced by {@link SNMPMIBParser#syntaxdesctype}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitSyntaxdesctype(SNMPMIBParser.SyntaxdesctypeContext ctx);
	/**
	 * Visit a parse tree produced by {@link SNMPMIBParser#indexdesctype}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitIndexdesctype(SNMPMIBParser.IndexdesctypeContext ctx);
	/**
	 * Visit a parse tree produced by {@link SNMPMIBParser#unknowndesctype}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitUnknowndesctype(SNMPMIBParser.UnknowndesctypeContext ctx);
	/**
	 * Visit a parse tree produced by {@link SNMPMIBParser#octetStringType}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitOctetStringType(SNMPMIBParser.OctetStringTypeContext ctx);
	/**
	 * Visit a parse tree produced by {@link SNMPMIBParser#objectidentifiertype}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitObjectidentifiertype(SNMPMIBParser.ObjectidentifiertypeContext ctx);
	/**
	 * Visit a parse tree produced by {@link SNMPMIBParser#sequencetype}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitSequencetype(SNMPMIBParser.SequencetypeContext ctx);
	/**
	 * Visit a parse tree produced by {@link SNMPMIBParser#typedefine}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitTypedefine(SNMPMIBParser.TypedefineContext ctx);
	/**
	 * Visit a parse tree produced by {@link SNMPMIBParser#integerType}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitIntegerType(SNMPMIBParser.IntegerTypeContext ctx);
	/**
	 * Visit a parse tree produced by {@link SNMPMIBParser#namedNumberList}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitNamedNumberList(SNMPMIBParser.NamedNumberListContext ctx);
	/**
	 * Visit a parse tree produced by {@link SNMPMIBParser#namedNumber}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitNamedNumber(SNMPMIBParser.NamedNumberContext ctx);
	/**
	 * Visit a parse tree produced by {@link SNMPMIBParser#signedNumber}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitSignedNumber(SNMPMIBParser.SignedNumberContext ctx);
	/**
	 * Visit a parse tree produced by {@link SNMPMIBParser#sizeConstraint}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitSizeConstraint(SNMPMIBParser.SizeConstraintContext ctx);
	/**
	 * Visit a parse tree produced by {@link SNMPMIBParser#constraint}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitConstraint(SNMPMIBParser.ConstraintContext ctx);
	/**
	 * Visit a parse tree produced by {@link SNMPMIBParser#rangeconstraint}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitRangeconstraint(SNMPMIBParser.RangeconstraintContext ctx);
}