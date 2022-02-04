// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package zeta

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// InitializeMarketIndexes is the `initializeMarketIndexes` instruction.
type InitializeMarketIndexes struct {
	Nonce *uint8

	// [0] = [] state
	//
	// [1] = [WRITE] marketIndexes
	//
	// [2] = [WRITE, SIGNER] admin
	//
	// [3] = [] systemProgram
	//
	// [4] = [] zetaGroup
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewInitializeMarketIndexesInstructionBuilder creates a new `InitializeMarketIndexes` instruction builder.
func NewInitializeMarketIndexesInstructionBuilder() *InitializeMarketIndexes {
	nd := &InitializeMarketIndexes{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 5),
	}
	return nd
}

// SetNonce sets the "nonce" parameter.
func (inst *InitializeMarketIndexes) SetNonce(nonce uint8) *InitializeMarketIndexes {
	inst.Nonce = &nonce
	return inst
}

// SetStateAccount sets the "state" account.
func (inst *InitializeMarketIndexes) SetStateAccount(state ag_solanago.PublicKey) *InitializeMarketIndexes {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(state)
	return inst
}

// GetStateAccount gets the "state" account.
func (inst *InitializeMarketIndexes) GetStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetMarketIndexesAccount sets the "marketIndexes" account.
func (inst *InitializeMarketIndexes) SetMarketIndexesAccount(marketIndexes ag_solanago.PublicKey) *InitializeMarketIndexes {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(marketIndexes).WRITE()
	return inst
}

// GetMarketIndexesAccount gets the "marketIndexes" account.
func (inst *InitializeMarketIndexes) GetMarketIndexesAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetAdminAccount sets the "admin" account.
func (inst *InitializeMarketIndexes) SetAdminAccount(admin ag_solanago.PublicKey) *InitializeMarketIndexes {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(admin).WRITE().SIGNER()
	return inst
}

// GetAdminAccount gets the "admin" account.
func (inst *InitializeMarketIndexes) GetAdminAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *InitializeMarketIndexes) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *InitializeMarketIndexes {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *InitializeMarketIndexes) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetZetaGroupAccount sets the "zetaGroup" account.
func (inst *InitializeMarketIndexes) SetZetaGroupAccount(zetaGroup ag_solanago.PublicKey) *InitializeMarketIndexes {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(zetaGroup)
	return inst
}

// GetZetaGroupAccount gets the "zetaGroup" account.
func (inst *InitializeMarketIndexes) GetZetaGroupAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

func (inst InitializeMarketIndexes) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_InitializeMarketIndexes,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst InitializeMarketIndexes) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *InitializeMarketIndexes) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Nonce == nil {
			return errors.New("Nonce parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.State is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.MarketIndexes is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Admin is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.ZetaGroup is not set")
		}
	}
	return nil
}

func (inst *InitializeMarketIndexes) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("InitializeMarketIndexes")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Nonce", *inst.Nonce))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=5]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("        state", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("marketIndexes", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("        admin", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("systemProgram", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("    zetaGroup", inst.AccountMetaSlice.Get(4)))
					})
				})
		})
}

func (obj InitializeMarketIndexes) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Nonce` param:
	err = encoder.Encode(obj.Nonce)
	if err != nil {
		return err
	}
	return nil
}
func (obj *InitializeMarketIndexes) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Nonce`:
	err = decoder.Decode(&obj.Nonce)
	if err != nil {
		return err
	}
	return nil
}

// NewInitializeMarketIndexesInstruction declares a new InitializeMarketIndexes instruction with the provided parameters and accounts.
func NewInitializeMarketIndexesInstruction(
	// Parameters:
	nonce uint8,
	// Accounts:
	state ag_solanago.PublicKey,
	marketIndexes ag_solanago.PublicKey,
	admin ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	zetaGroup ag_solanago.PublicKey) *InitializeMarketIndexes {
	return NewInitializeMarketIndexesInstructionBuilder().
		SetNonce(nonce).
		SetStateAccount(state).
		SetMarketIndexesAccount(marketIndexes).
		SetAdminAccount(admin).
		SetSystemProgramAccount(systemProgram).
		SetZetaGroupAccount(zetaGroup)
}
