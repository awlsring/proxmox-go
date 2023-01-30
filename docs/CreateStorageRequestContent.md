# CreateStorageRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Storage** | **string** |  | 
**Type** | [**StorageType**](StorageType.md) |  | 
**Authsupported** | Pointer to **string** |  | [optional] 
**Base** | Pointer to **string** | The base volume to use.  | [optional] 
**Blocksize** | Pointer to **string** |  | [optional] 
**Bwlimit** | Pointer to **string** |  | [optional] 
**ComstarHg** | Pointer to **string** |  | [optional] 
**ComstarTg** | Pointer to **string** |  | [optional] 
**Content** | Pointer to **string** | Allowed content types. | [optional] 
**DataPool** | Pointer to **string** | Data pool, only for erasure coding | [optional] 
**Disable** | Pointer to **bool** | Disable the storage | [optional] 
**Domain** | Pointer to **string** | CIFS domain | [optional] 
**EncryptionKey** | Pointer to **string** | Encryption key for storage. Use &#39;autogen&#39; to generate one automatically without passphrase. | [optional] 
**Export** | Pointer to **string** | NFS export path | [optional] 
**Fingerprint** | Pointer to **string** | Certificate SHA256 fingerprint | [optional] 
**Format** | Pointer to **string** | Default image format | [optional] 
**FsName** | Pointer to **string** | The ceph filesystem name | [optional] 
**Fuse** | Pointer to **bool** | Mount CephFS through FUSE | [optional] 
**IsMountpoint** | Pointer to **string** | Assume the given path is an externally managed mountpoint. This is a string boolean, using &#39;yes&#39; or &#39;no&#39; | [optional] 
**Iscsiprovider** | Pointer to **string** | iSCSI provider | [optional] 
**Keyring** | Pointer to **string** | Client keyring contents | [optional] 
**Krbd** | Pointer to **bool** | Always access rbd through krbd kernel module | [optional] 
**LioTpg** | Pointer to **string** | Target portal group for Linux LIO targets | [optional] 
**MasterPubkey** | Pointer to **string** | Base64 encoded, PEM formatted public RSA key | [optional] 
**MaxProtectedBackups** | Pointer to **float32** | Mac number of protected backups. Defaults to unlimited | [optional] 
**Mkdir** | Pointer to **bool** | Create the directory if doesnt exist. Defaults to true. | [optional] 
**Monhost** | Pointer to **string** | IP addresses of monitors. | [optional] 
**Mountpoint** | Pointer to **string** | The mount point | [optional] 
**Namespace** | Pointer to **string** | The namespace | [optional] 
**Nocow** | Pointer to **float32** | Set the NOCOW flag on files. This is a boolean integer, using 0 for false or 1 for true. | [optional] 
**Nodes** | Pointer to **string** | Comma seperated list of node names | [optional] 
**Nowritecache** | Pointer to **bool** | Disable write caching on the target. | [optional] 
**Options** | Pointer to **string** | NFS mount options. | [optional] 
**Password** | Pointer to **string** | Password for accessing the datastore | [optional] 
**Path** | Pointer to **string** | The filesystem path | [optional] 
**Pool** | Pointer to **string** | The pool name | [optional] 
**Port** | Pointer to **float32** | For non default port | [optional] 
**Portal** | Pointer to **string** | iSCSI portal (IP or DNS name with optional port). | [optional] 
**Preallocation** | Pointer to [**StoragePreallocation**](StoragePreallocation.md) |  | [optional] 
**PruneBackups** | Pointer to **string** | The retention options. | [optional] 
**Saferemove** | Pointer to **bool** | Zero out data when removing LVs | [optional] 
**SaferemoveThroughput** | Pointer to **string** | Wipe throughput. | [optional] 
**Server** | Pointer to **string** | Server IP or DNS name | [optional] 
**Server2** | Pointer to **string** | Backup server IP or DNS name | [optional] 
**Share** | Pointer to **string** | The CIFS share. | [optional] 
**Shared** | Pointer to **bool** | Indicates if the storage is shared. | [optional] 
**Smbversion** | Pointer to [**StorageSMBVersion**](StorageSMBVersion.md) |  | [optional] 
**Sparse** | Pointer to **bool** | Use sparse volumes. | [optional] 
**Subdir** | Pointer to **string** | The subdir to mount. | [optional] 
**TaggedOnly** | Pointer to **bool** | Only use logical volumes tagged with &#39;pve-vm-ID&#39;. | [optional] 
**Thinpool** | Pointer to **string** | LVM thin pool name. | [optional] 
**Transport** | Pointer to [**StorageTransport**](StorageTransport.md) |  | [optional] 
**Username** | Pointer to **string** | The username to use. | [optional] 
**Vgname** | Pointer to **string** | The volume group name. | [optional] 
**Volume** | Pointer to **string** | The glusterfs volume. | [optional] 

## Methods

### NewCreateStorageRequestContent

`func NewCreateStorageRequestContent(storage string, type_ StorageType, ) *CreateStorageRequestContent`

NewCreateStorageRequestContent instantiates a new CreateStorageRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateStorageRequestContentWithDefaults

`func NewCreateStorageRequestContentWithDefaults() *CreateStorageRequestContent`

NewCreateStorageRequestContentWithDefaults instantiates a new CreateStorageRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStorage

`func (o *CreateStorageRequestContent) GetStorage() string`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *CreateStorageRequestContent) GetStorageOk() (*string, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *CreateStorageRequestContent) SetStorage(v string)`

SetStorage sets Storage field to given value.


### GetType

`func (o *CreateStorageRequestContent) GetType() StorageType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateStorageRequestContent) GetTypeOk() (*StorageType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateStorageRequestContent) SetType(v StorageType)`

SetType sets Type field to given value.


### GetAuthsupported

`func (o *CreateStorageRequestContent) GetAuthsupported() string`

GetAuthsupported returns the Authsupported field if non-nil, zero value otherwise.

### GetAuthsupportedOk

`func (o *CreateStorageRequestContent) GetAuthsupportedOk() (*string, bool)`

GetAuthsupportedOk returns a tuple with the Authsupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthsupported

`func (o *CreateStorageRequestContent) SetAuthsupported(v string)`

SetAuthsupported sets Authsupported field to given value.

### HasAuthsupported

`func (o *CreateStorageRequestContent) HasAuthsupported() bool`

HasAuthsupported returns a boolean if a field has been set.

### GetBase

`func (o *CreateStorageRequestContent) GetBase() string`

GetBase returns the Base field if non-nil, zero value otherwise.

### GetBaseOk

`func (o *CreateStorageRequestContent) GetBaseOk() (*string, bool)`

GetBaseOk returns a tuple with the Base field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase

`func (o *CreateStorageRequestContent) SetBase(v string)`

SetBase sets Base field to given value.

### HasBase

`func (o *CreateStorageRequestContent) HasBase() bool`

HasBase returns a boolean if a field has been set.

### GetBlocksize

`func (o *CreateStorageRequestContent) GetBlocksize() string`

GetBlocksize returns the Blocksize field if non-nil, zero value otherwise.

### GetBlocksizeOk

`func (o *CreateStorageRequestContent) GetBlocksizeOk() (*string, bool)`

GetBlocksizeOk returns a tuple with the Blocksize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocksize

`func (o *CreateStorageRequestContent) SetBlocksize(v string)`

SetBlocksize sets Blocksize field to given value.

### HasBlocksize

`func (o *CreateStorageRequestContent) HasBlocksize() bool`

HasBlocksize returns a boolean if a field has been set.

### GetBwlimit

`func (o *CreateStorageRequestContent) GetBwlimit() string`

GetBwlimit returns the Bwlimit field if non-nil, zero value otherwise.

### GetBwlimitOk

`func (o *CreateStorageRequestContent) GetBwlimitOk() (*string, bool)`

GetBwlimitOk returns a tuple with the Bwlimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBwlimit

`func (o *CreateStorageRequestContent) SetBwlimit(v string)`

SetBwlimit sets Bwlimit field to given value.

### HasBwlimit

`func (o *CreateStorageRequestContent) HasBwlimit() bool`

HasBwlimit returns a boolean if a field has been set.

### GetComstarHg

`func (o *CreateStorageRequestContent) GetComstarHg() string`

GetComstarHg returns the ComstarHg field if non-nil, zero value otherwise.

### GetComstarHgOk

`func (o *CreateStorageRequestContent) GetComstarHgOk() (*string, bool)`

GetComstarHgOk returns a tuple with the ComstarHg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComstarHg

`func (o *CreateStorageRequestContent) SetComstarHg(v string)`

SetComstarHg sets ComstarHg field to given value.

### HasComstarHg

`func (o *CreateStorageRequestContent) HasComstarHg() bool`

HasComstarHg returns a boolean if a field has been set.

### GetComstarTg

`func (o *CreateStorageRequestContent) GetComstarTg() string`

GetComstarTg returns the ComstarTg field if non-nil, zero value otherwise.

### GetComstarTgOk

`func (o *CreateStorageRequestContent) GetComstarTgOk() (*string, bool)`

GetComstarTgOk returns a tuple with the ComstarTg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComstarTg

`func (o *CreateStorageRequestContent) SetComstarTg(v string)`

SetComstarTg sets ComstarTg field to given value.

### HasComstarTg

`func (o *CreateStorageRequestContent) HasComstarTg() bool`

HasComstarTg returns a boolean if a field has been set.

### GetContent

`func (o *CreateStorageRequestContent) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *CreateStorageRequestContent) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *CreateStorageRequestContent) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *CreateStorageRequestContent) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetDataPool

`func (o *CreateStorageRequestContent) GetDataPool() string`

GetDataPool returns the DataPool field if non-nil, zero value otherwise.

### GetDataPoolOk

`func (o *CreateStorageRequestContent) GetDataPoolOk() (*string, bool)`

GetDataPoolOk returns a tuple with the DataPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataPool

`func (o *CreateStorageRequestContent) SetDataPool(v string)`

SetDataPool sets DataPool field to given value.

### HasDataPool

`func (o *CreateStorageRequestContent) HasDataPool() bool`

HasDataPool returns a boolean if a field has been set.

### GetDisable

`func (o *CreateStorageRequestContent) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *CreateStorageRequestContent) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *CreateStorageRequestContent) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *CreateStorageRequestContent) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetDomain

`func (o *CreateStorageRequestContent) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *CreateStorageRequestContent) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *CreateStorageRequestContent) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *CreateStorageRequestContent) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetEncryptionKey

`func (o *CreateStorageRequestContent) GetEncryptionKey() string`

GetEncryptionKey returns the EncryptionKey field if non-nil, zero value otherwise.

### GetEncryptionKeyOk

`func (o *CreateStorageRequestContent) GetEncryptionKeyOk() (*string, bool)`

GetEncryptionKeyOk returns a tuple with the EncryptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionKey

`func (o *CreateStorageRequestContent) SetEncryptionKey(v string)`

SetEncryptionKey sets EncryptionKey field to given value.

### HasEncryptionKey

`func (o *CreateStorageRequestContent) HasEncryptionKey() bool`

HasEncryptionKey returns a boolean if a field has been set.

### GetExport

`func (o *CreateStorageRequestContent) GetExport() string`

GetExport returns the Export field if non-nil, zero value otherwise.

### GetExportOk

`func (o *CreateStorageRequestContent) GetExportOk() (*string, bool)`

GetExportOk returns a tuple with the Export field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExport

`func (o *CreateStorageRequestContent) SetExport(v string)`

SetExport sets Export field to given value.

### HasExport

`func (o *CreateStorageRequestContent) HasExport() bool`

HasExport returns a boolean if a field has been set.

### GetFingerprint

`func (o *CreateStorageRequestContent) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *CreateStorageRequestContent) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *CreateStorageRequestContent) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *CreateStorageRequestContent) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### GetFormat

`func (o *CreateStorageRequestContent) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *CreateStorageRequestContent) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *CreateStorageRequestContent) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *CreateStorageRequestContent) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetFsName

`func (o *CreateStorageRequestContent) GetFsName() string`

GetFsName returns the FsName field if non-nil, zero value otherwise.

### GetFsNameOk

`func (o *CreateStorageRequestContent) GetFsNameOk() (*string, bool)`

GetFsNameOk returns a tuple with the FsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsName

`func (o *CreateStorageRequestContent) SetFsName(v string)`

SetFsName sets FsName field to given value.

### HasFsName

`func (o *CreateStorageRequestContent) HasFsName() bool`

HasFsName returns a boolean if a field has been set.

### GetFuse

`func (o *CreateStorageRequestContent) GetFuse() bool`

GetFuse returns the Fuse field if non-nil, zero value otherwise.

### GetFuseOk

`func (o *CreateStorageRequestContent) GetFuseOk() (*bool, bool)`

GetFuseOk returns a tuple with the Fuse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuse

`func (o *CreateStorageRequestContent) SetFuse(v bool)`

SetFuse sets Fuse field to given value.

### HasFuse

`func (o *CreateStorageRequestContent) HasFuse() bool`

HasFuse returns a boolean if a field has been set.

### GetIsMountpoint

`func (o *CreateStorageRequestContent) GetIsMountpoint() string`

GetIsMountpoint returns the IsMountpoint field if non-nil, zero value otherwise.

### GetIsMountpointOk

`func (o *CreateStorageRequestContent) GetIsMountpointOk() (*string, bool)`

GetIsMountpointOk returns a tuple with the IsMountpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMountpoint

`func (o *CreateStorageRequestContent) SetIsMountpoint(v string)`

SetIsMountpoint sets IsMountpoint field to given value.

### HasIsMountpoint

`func (o *CreateStorageRequestContent) HasIsMountpoint() bool`

HasIsMountpoint returns a boolean if a field has been set.

### GetIscsiprovider

`func (o *CreateStorageRequestContent) GetIscsiprovider() string`

GetIscsiprovider returns the Iscsiprovider field if non-nil, zero value otherwise.

### GetIscsiproviderOk

`func (o *CreateStorageRequestContent) GetIscsiproviderOk() (*string, bool)`

GetIscsiproviderOk returns a tuple with the Iscsiprovider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIscsiprovider

`func (o *CreateStorageRequestContent) SetIscsiprovider(v string)`

SetIscsiprovider sets Iscsiprovider field to given value.

### HasIscsiprovider

`func (o *CreateStorageRequestContent) HasIscsiprovider() bool`

HasIscsiprovider returns a boolean if a field has been set.

### GetKeyring

`func (o *CreateStorageRequestContent) GetKeyring() string`

GetKeyring returns the Keyring field if non-nil, zero value otherwise.

### GetKeyringOk

`func (o *CreateStorageRequestContent) GetKeyringOk() (*string, bool)`

GetKeyringOk returns a tuple with the Keyring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyring

`func (o *CreateStorageRequestContent) SetKeyring(v string)`

SetKeyring sets Keyring field to given value.

### HasKeyring

`func (o *CreateStorageRequestContent) HasKeyring() bool`

HasKeyring returns a boolean if a field has been set.

### GetKrbd

`func (o *CreateStorageRequestContent) GetKrbd() bool`

GetKrbd returns the Krbd field if non-nil, zero value otherwise.

### GetKrbdOk

`func (o *CreateStorageRequestContent) GetKrbdOk() (*bool, bool)`

GetKrbdOk returns a tuple with the Krbd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKrbd

`func (o *CreateStorageRequestContent) SetKrbd(v bool)`

SetKrbd sets Krbd field to given value.

### HasKrbd

`func (o *CreateStorageRequestContent) HasKrbd() bool`

HasKrbd returns a boolean if a field has been set.

### GetLioTpg

`func (o *CreateStorageRequestContent) GetLioTpg() string`

GetLioTpg returns the LioTpg field if non-nil, zero value otherwise.

### GetLioTpgOk

`func (o *CreateStorageRequestContent) GetLioTpgOk() (*string, bool)`

GetLioTpgOk returns a tuple with the LioTpg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLioTpg

`func (o *CreateStorageRequestContent) SetLioTpg(v string)`

SetLioTpg sets LioTpg field to given value.

### HasLioTpg

`func (o *CreateStorageRequestContent) HasLioTpg() bool`

HasLioTpg returns a boolean if a field has been set.

### GetMasterPubkey

`func (o *CreateStorageRequestContent) GetMasterPubkey() string`

GetMasterPubkey returns the MasterPubkey field if non-nil, zero value otherwise.

### GetMasterPubkeyOk

`func (o *CreateStorageRequestContent) GetMasterPubkeyOk() (*string, bool)`

GetMasterPubkeyOk returns a tuple with the MasterPubkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterPubkey

`func (o *CreateStorageRequestContent) SetMasterPubkey(v string)`

SetMasterPubkey sets MasterPubkey field to given value.

### HasMasterPubkey

`func (o *CreateStorageRequestContent) HasMasterPubkey() bool`

HasMasterPubkey returns a boolean if a field has been set.

### GetMaxProtectedBackups

`func (o *CreateStorageRequestContent) GetMaxProtectedBackups() float32`

GetMaxProtectedBackups returns the MaxProtectedBackups field if non-nil, zero value otherwise.

### GetMaxProtectedBackupsOk

`func (o *CreateStorageRequestContent) GetMaxProtectedBackupsOk() (*float32, bool)`

GetMaxProtectedBackupsOk returns a tuple with the MaxProtectedBackups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxProtectedBackups

`func (o *CreateStorageRequestContent) SetMaxProtectedBackups(v float32)`

SetMaxProtectedBackups sets MaxProtectedBackups field to given value.

### HasMaxProtectedBackups

`func (o *CreateStorageRequestContent) HasMaxProtectedBackups() bool`

HasMaxProtectedBackups returns a boolean if a field has been set.

### GetMkdir

`func (o *CreateStorageRequestContent) GetMkdir() bool`

GetMkdir returns the Mkdir field if non-nil, zero value otherwise.

### GetMkdirOk

`func (o *CreateStorageRequestContent) GetMkdirOk() (*bool, bool)`

GetMkdirOk returns a tuple with the Mkdir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMkdir

`func (o *CreateStorageRequestContent) SetMkdir(v bool)`

SetMkdir sets Mkdir field to given value.

### HasMkdir

`func (o *CreateStorageRequestContent) HasMkdir() bool`

HasMkdir returns a boolean if a field has been set.

### GetMonhost

`func (o *CreateStorageRequestContent) GetMonhost() string`

GetMonhost returns the Monhost field if non-nil, zero value otherwise.

### GetMonhostOk

`func (o *CreateStorageRequestContent) GetMonhostOk() (*string, bool)`

GetMonhostOk returns a tuple with the Monhost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonhost

`func (o *CreateStorageRequestContent) SetMonhost(v string)`

SetMonhost sets Monhost field to given value.

### HasMonhost

`func (o *CreateStorageRequestContent) HasMonhost() bool`

HasMonhost returns a boolean if a field has been set.

### GetMountpoint

`func (o *CreateStorageRequestContent) GetMountpoint() string`

GetMountpoint returns the Mountpoint field if non-nil, zero value otherwise.

### GetMountpointOk

`func (o *CreateStorageRequestContent) GetMountpointOk() (*string, bool)`

GetMountpointOk returns a tuple with the Mountpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountpoint

`func (o *CreateStorageRequestContent) SetMountpoint(v string)`

SetMountpoint sets Mountpoint field to given value.

### HasMountpoint

`func (o *CreateStorageRequestContent) HasMountpoint() bool`

HasMountpoint returns a boolean if a field has been set.

### GetNamespace

`func (o *CreateStorageRequestContent) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *CreateStorageRequestContent) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *CreateStorageRequestContent) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *CreateStorageRequestContent) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetNocow

`func (o *CreateStorageRequestContent) GetNocow() float32`

GetNocow returns the Nocow field if non-nil, zero value otherwise.

### GetNocowOk

`func (o *CreateStorageRequestContent) GetNocowOk() (*float32, bool)`

GetNocowOk returns a tuple with the Nocow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNocow

`func (o *CreateStorageRequestContent) SetNocow(v float32)`

SetNocow sets Nocow field to given value.

### HasNocow

`func (o *CreateStorageRequestContent) HasNocow() bool`

HasNocow returns a boolean if a field has been set.

### GetNodes

`func (o *CreateStorageRequestContent) GetNodes() string`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *CreateStorageRequestContent) GetNodesOk() (*string, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *CreateStorageRequestContent) SetNodes(v string)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *CreateStorageRequestContent) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetNowritecache

`func (o *CreateStorageRequestContent) GetNowritecache() bool`

GetNowritecache returns the Nowritecache field if non-nil, zero value otherwise.

### GetNowritecacheOk

`func (o *CreateStorageRequestContent) GetNowritecacheOk() (*bool, bool)`

GetNowritecacheOk returns a tuple with the Nowritecache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNowritecache

`func (o *CreateStorageRequestContent) SetNowritecache(v bool)`

SetNowritecache sets Nowritecache field to given value.

### HasNowritecache

`func (o *CreateStorageRequestContent) HasNowritecache() bool`

HasNowritecache returns a boolean if a field has been set.

### GetOptions

`func (o *CreateStorageRequestContent) GetOptions() string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *CreateStorageRequestContent) GetOptionsOk() (*string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *CreateStorageRequestContent) SetOptions(v string)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *CreateStorageRequestContent) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetPassword

`func (o *CreateStorageRequestContent) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateStorageRequestContent) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateStorageRequestContent) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CreateStorageRequestContent) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPath

`func (o *CreateStorageRequestContent) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *CreateStorageRequestContent) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *CreateStorageRequestContent) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *CreateStorageRequestContent) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetPool

`func (o *CreateStorageRequestContent) GetPool() string`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *CreateStorageRequestContent) GetPoolOk() (*string, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *CreateStorageRequestContent) SetPool(v string)`

SetPool sets Pool field to given value.

### HasPool

`func (o *CreateStorageRequestContent) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetPort

`func (o *CreateStorageRequestContent) GetPort() float32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *CreateStorageRequestContent) GetPortOk() (*float32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *CreateStorageRequestContent) SetPort(v float32)`

SetPort sets Port field to given value.

### HasPort

`func (o *CreateStorageRequestContent) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetPortal

`func (o *CreateStorageRequestContent) GetPortal() string`

GetPortal returns the Portal field if non-nil, zero value otherwise.

### GetPortalOk

`func (o *CreateStorageRequestContent) GetPortalOk() (*string, bool)`

GetPortalOk returns a tuple with the Portal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortal

`func (o *CreateStorageRequestContent) SetPortal(v string)`

SetPortal sets Portal field to given value.

### HasPortal

`func (o *CreateStorageRequestContent) HasPortal() bool`

HasPortal returns a boolean if a field has been set.

### GetPreallocation

`func (o *CreateStorageRequestContent) GetPreallocation() StoragePreallocation`

GetPreallocation returns the Preallocation field if non-nil, zero value otherwise.

### GetPreallocationOk

`func (o *CreateStorageRequestContent) GetPreallocationOk() (*StoragePreallocation, bool)`

GetPreallocationOk returns a tuple with the Preallocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreallocation

`func (o *CreateStorageRequestContent) SetPreallocation(v StoragePreallocation)`

SetPreallocation sets Preallocation field to given value.

### HasPreallocation

`func (o *CreateStorageRequestContent) HasPreallocation() bool`

HasPreallocation returns a boolean if a field has been set.

### GetPruneBackups

`func (o *CreateStorageRequestContent) GetPruneBackups() string`

GetPruneBackups returns the PruneBackups field if non-nil, zero value otherwise.

### GetPruneBackupsOk

`func (o *CreateStorageRequestContent) GetPruneBackupsOk() (*string, bool)`

GetPruneBackupsOk returns a tuple with the PruneBackups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPruneBackups

`func (o *CreateStorageRequestContent) SetPruneBackups(v string)`

SetPruneBackups sets PruneBackups field to given value.

### HasPruneBackups

`func (o *CreateStorageRequestContent) HasPruneBackups() bool`

HasPruneBackups returns a boolean if a field has been set.

### GetSaferemove

`func (o *CreateStorageRequestContent) GetSaferemove() bool`

GetSaferemove returns the Saferemove field if non-nil, zero value otherwise.

### GetSaferemoveOk

`func (o *CreateStorageRequestContent) GetSaferemoveOk() (*bool, bool)`

GetSaferemoveOk returns a tuple with the Saferemove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaferemove

`func (o *CreateStorageRequestContent) SetSaferemove(v bool)`

SetSaferemove sets Saferemove field to given value.

### HasSaferemove

`func (o *CreateStorageRequestContent) HasSaferemove() bool`

HasSaferemove returns a boolean if a field has been set.

### GetSaferemoveThroughput

`func (o *CreateStorageRequestContent) GetSaferemoveThroughput() string`

GetSaferemoveThroughput returns the SaferemoveThroughput field if non-nil, zero value otherwise.

### GetSaferemoveThroughputOk

`func (o *CreateStorageRequestContent) GetSaferemoveThroughputOk() (*string, bool)`

GetSaferemoveThroughputOk returns a tuple with the SaferemoveThroughput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaferemoveThroughput

`func (o *CreateStorageRequestContent) SetSaferemoveThroughput(v string)`

SetSaferemoveThroughput sets SaferemoveThroughput field to given value.

### HasSaferemoveThroughput

`func (o *CreateStorageRequestContent) HasSaferemoveThroughput() bool`

HasSaferemoveThroughput returns a boolean if a field has been set.

### GetServer

`func (o *CreateStorageRequestContent) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *CreateStorageRequestContent) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *CreateStorageRequestContent) SetServer(v string)`

SetServer sets Server field to given value.

### HasServer

`func (o *CreateStorageRequestContent) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetServer2

`func (o *CreateStorageRequestContent) GetServer2() string`

GetServer2 returns the Server2 field if non-nil, zero value otherwise.

### GetServer2Ok

`func (o *CreateStorageRequestContent) GetServer2Ok() (*string, bool)`

GetServer2Ok returns a tuple with the Server2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer2

`func (o *CreateStorageRequestContent) SetServer2(v string)`

SetServer2 sets Server2 field to given value.

### HasServer2

`func (o *CreateStorageRequestContent) HasServer2() bool`

HasServer2 returns a boolean if a field has been set.

### GetShare

`func (o *CreateStorageRequestContent) GetShare() string`

GetShare returns the Share field if non-nil, zero value otherwise.

### GetShareOk

`func (o *CreateStorageRequestContent) GetShareOk() (*string, bool)`

GetShareOk returns a tuple with the Share field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShare

`func (o *CreateStorageRequestContent) SetShare(v string)`

SetShare sets Share field to given value.

### HasShare

`func (o *CreateStorageRequestContent) HasShare() bool`

HasShare returns a boolean if a field has been set.

### GetShared

`func (o *CreateStorageRequestContent) GetShared() bool`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *CreateStorageRequestContent) GetSharedOk() (*bool, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *CreateStorageRequestContent) SetShared(v bool)`

SetShared sets Shared field to given value.

### HasShared

`func (o *CreateStorageRequestContent) HasShared() bool`

HasShared returns a boolean if a field has been set.

### GetSmbversion

`func (o *CreateStorageRequestContent) GetSmbversion() StorageSMBVersion`

GetSmbversion returns the Smbversion field if non-nil, zero value otherwise.

### GetSmbversionOk

`func (o *CreateStorageRequestContent) GetSmbversionOk() (*StorageSMBVersion, bool)`

GetSmbversionOk returns a tuple with the Smbversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmbversion

`func (o *CreateStorageRequestContent) SetSmbversion(v StorageSMBVersion)`

SetSmbversion sets Smbversion field to given value.

### HasSmbversion

`func (o *CreateStorageRequestContent) HasSmbversion() bool`

HasSmbversion returns a boolean if a field has been set.

### GetSparse

`func (o *CreateStorageRequestContent) GetSparse() bool`

GetSparse returns the Sparse field if non-nil, zero value otherwise.

### GetSparseOk

`func (o *CreateStorageRequestContent) GetSparseOk() (*bool, bool)`

GetSparseOk returns a tuple with the Sparse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSparse

`func (o *CreateStorageRequestContent) SetSparse(v bool)`

SetSparse sets Sparse field to given value.

### HasSparse

`func (o *CreateStorageRequestContent) HasSparse() bool`

HasSparse returns a boolean if a field has been set.

### GetSubdir

`func (o *CreateStorageRequestContent) GetSubdir() string`

GetSubdir returns the Subdir field if non-nil, zero value otherwise.

### GetSubdirOk

`func (o *CreateStorageRequestContent) GetSubdirOk() (*string, bool)`

GetSubdirOk returns a tuple with the Subdir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdir

`func (o *CreateStorageRequestContent) SetSubdir(v string)`

SetSubdir sets Subdir field to given value.

### HasSubdir

`func (o *CreateStorageRequestContent) HasSubdir() bool`

HasSubdir returns a boolean if a field has been set.

### GetTaggedOnly

`func (o *CreateStorageRequestContent) GetTaggedOnly() bool`

GetTaggedOnly returns the TaggedOnly field if non-nil, zero value otherwise.

### GetTaggedOnlyOk

`func (o *CreateStorageRequestContent) GetTaggedOnlyOk() (*bool, bool)`

GetTaggedOnlyOk returns a tuple with the TaggedOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaggedOnly

`func (o *CreateStorageRequestContent) SetTaggedOnly(v bool)`

SetTaggedOnly sets TaggedOnly field to given value.

### HasTaggedOnly

`func (o *CreateStorageRequestContent) HasTaggedOnly() bool`

HasTaggedOnly returns a boolean if a field has been set.

### GetThinpool

`func (o *CreateStorageRequestContent) GetThinpool() string`

GetThinpool returns the Thinpool field if non-nil, zero value otherwise.

### GetThinpoolOk

`func (o *CreateStorageRequestContent) GetThinpoolOk() (*string, bool)`

GetThinpoolOk returns a tuple with the Thinpool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThinpool

`func (o *CreateStorageRequestContent) SetThinpool(v string)`

SetThinpool sets Thinpool field to given value.

### HasThinpool

`func (o *CreateStorageRequestContent) HasThinpool() bool`

HasThinpool returns a boolean if a field has been set.

### GetTransport

`func (o *CreateStorageRequestContent) GetTransport() StorageTransport`

GetTransport returns the Transport field if non-nil, zero value otherwise.

### GetTransportOk

`func (o *CreateStorageRequestContent) GetTransportOk() (*StorageTransport, bool)`

GetTransportOk returns a tuple with the Transport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransport

`func (o *CreateStorageRequestContent) SetTransport(v StorageTransport)`

SetTransport sets Transport field to given value.

### HasTransport

`func (o *CreateStorageRequestContent) HasTransport() bool`

HasTransport returns a boolean if a field has been set.

### GetUsername

`func (o *CreateStorageRequestContent) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CreateStorageRequestContent) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CreateStorageRequestContent) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *CreateStorageRequestContent) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetVgname

`func (o *CreateStorageRequestContent) GetVgname() string`

GetVgname returns the Vgname field if non-nil, zero value otherwise.

### GetVgnameOk

`func (o *CreateStorageRequestContent) GetVgnameOk() (*string, bool)`

GetVgnameOk returns a tuple with the Vgname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVgname

`func (o *CreateStorageRequestContent) SetVgname(v string)`

SetVgname sets Vgname field to given value.

### HasVgname

`func (o *CreateStorageRequestContent) HasVgname() bool`

HasVgname returns a boolean if a field has been set.

### GetVolume

`func (o *CreateStorageRequestContent) GetVolume() string`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *CreateStorageRequestContent) GetVolumeOk() (*string, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *CreateStorageRequestContent) SetVolume(v string)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *CreateStorageRequestContent) HasVolume() bool`

HasVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


