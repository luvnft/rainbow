// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package zeta

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// BurnVaultTokens is the `burnVaultTokens` instruction.
type BurnVaultTokens struct {

	// [0] = [] state
	//
	// [1] = [WRITE] mint
	//
	// [2] = [WRITE] vault
	//
	// [3] = [] serumAuthority
	//
	// [4] = [] tokenProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewBurnVaultTokensInstructionBuilder creates a new `BurnVaultTokens` instruction builder.
func NewBurnVaultTokensInstructionBuilder() *BurnVaultTokens {
	nd := &BurnVaultTokens{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 5),
	}
	return nd
}

// SetStateAccount sets the "state" account.
func (inst *BurnVaultTokens) SetStateAccount(state ag_solanago.PublicKey) *BurnVaultTokens {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(state)
	return inst
}

// GetStateAccount gets the "state" account.
func (inst *BurnVaultTokens) GetStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetMintAccount sets the "mint" account.
func (inst *BurnVaultTokens) SetMintAccount(mint ag_solanago.PublicKey) *BurnVaultTokens {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(mint).WRITE()
	return inst
}

// GetMintAccount gets the "mint" account.
func (inst *BurnVaultTokens) GetMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetVaultAccount sets the "vault" account.
func (inst *BurnVaultTokens) SetVaultAccount(vault ag_solanago.PublicKey) *BurnVaultTokens {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(vault).WRITE()
	return inst
}

// GetVaultAccount gets the "vault" account.
func (inst *BurnVaultTokens) GetVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetSerumAuthorityAccount sets the "serumAuthority" account.
func (inst *BurnVaultTokens) SetSerumAuthorityAccount(serumAuthority ag_solanago.PublicKey) *BurnVaultTokens {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(serumAuthority)
	return inst
}

// GetSerumAuthorityAccount gets the "serumAuthority" account.
func (inst *BurnVaultTokens) GetSerumAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *BurnVaultTokens) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *BurnVaultTokens {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *BurnVaultTokens) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

func (inst BurnVaultTokens) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_BurnVaultTokens,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst BurnVaultTokens) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *BurnVaultTokens) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.State is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Mint is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Vault is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.SerumAuthority is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
	}
	return nil
}

func (inst *BurnVaultTokens) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("BurnVaultTokens")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=5]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("         state", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("          mint", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("         vault", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("serumAuthority", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("  tokenProgram", inst.AccountMetaSlice.Get(4)))
					})
				})
		})
}

func (obj BurnVaultTokens) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *BurnVaultTokens) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewBurnVaultTokensInstruction declares a new BurnVaultTokens instruction with the provided parameters and accounts.
func NewBurnVaultTokensInstruction(
	// Accounts:
	state ag_solanago.PublicKey,
	mint ag_solanago.PublicKey,
	vault ag_solanago.PublicKey,
	serumAuthority ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey) *BurnVaultTokens {
	return NewBurnVaultTokensInstructionBuilder().
		SetStateAccount(state).
		SetMintAccount(mint).
		SetVaultAccount(vault).
		SetSerumAuthorityAccount(serumAuthority).
		SetTokenProgramAccount(tokenProgram)
}